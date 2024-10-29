package main

import (
	"gim/cmd/logic/middleware"
	"gim/config"
	"gim/internal/logic/api"
	"gim/internal/logic/domain/device"
	"gim/internal/logic/domain/message"
	"gim/internal/logic/proxy"
	"gim/pkg/interceptor"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/urlwhitelist"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"gim/internal/logic/apisocket"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	proxy.MessageProxy = message.App
	proxy.DeviceProxy = device.App
}
func setRouter(r *gin.Engine) {
	v1 := r.Group("/im/")
	{
		v1.POST("/register_device", RegisterDevice)
		v1.POST("/gettoken", GetToken)
	}

}
func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.NewInterceptor("logic_interceptor", urlwhitelist.Logic)))

	// 监听服务关闭信号，服务平滑重启
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		s := <-c
		logger.Logger.Info("server stop", zap.Any("signal", s))
		server.GracefulStop()
	}()

	pb.RegisterLogicIntServer(server, &api.LogicIntServer{})
	pb.RegisterLogicExtServer(server, &api.LogicExtServer{})
	listen, err := net.Listen("tcp", config.Config.LogicRPCListenAddr)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.CrosHandler())
	setRouter(r)
	logger.Logger.Info("http端口启动在8080")
	go func() {
		r.Run(":8888")
	}()

	// 启动WebSocket长链接服务器
	go func() {
		wsPort := ":8112"
		apisocket.StartWSServer(wsPort)
		logger.Logger.Info("启动websocket端口" + wsPort)
	}()

	logger.Logger.Info("rpc服务已经开启")
	err = server.Serve(listen)
	if err != nil {
		logger.Logger.Error("serve error", zap.Error(err))
	}

}
