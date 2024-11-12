package config

import (
	"context"
	"gim/pkg/gerrors"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"go.uber.org/zap"
	"os"

	"google.golang.org/grpc"
)

var builders = map[string]Builder{
	"default": &defaultBuilder{},
	"k8s":     &k8sBuilder{},
	"online":  &defaultBuilderOnline{},
}

var Config Configuration

type Builder interface {
	Build() Configuration
}

type Configuration struct {
	MySQL                string
	RedisHost            string
	RedisPassword        string
	PushRoomSubscribeNum int
	PushAllSubscribeNum  int

	ConnectLocalAddr     string
	ConnectRPCListenAddr string
	ConnectTCPListenAddr string
	ConnectWSListenAddr  string

	LogicRPCListenAddr    string
	BusinessRPCListenAddr string
	FileHTTPListenAddr    string

	ConnectIntClientBuilder  func() pb.ConnectIntClient
	LogicIntClientBuilder    func() pb.LogicIntClient
	BusinessIntClientBuilder func() pb.BusinessIntClient

	BusinessExtClientBuilder func() pb.BusinessExtClient //@ms
	LogicExtClientBuilder    func() pb.LogicExtClient    //@ms
}

func init() {
	env := os.Getenv("GIM_ENV")
	logger.Logger.Info("init env :", zap.Any("env", env))
	builder, ok := builders[env]
	if !ok {
		builder = new(defaultBuilder)
	}
	Config = builder.Build()

}

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return gerrors.WrapRPCError(err)
}
