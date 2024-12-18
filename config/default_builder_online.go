package config

import (
	"context"
	"fmt"
	"gim/pkg/grpclib/picker"
	_ "gim/pkg/grpclib/resolver/addrs"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type defaultBuilderOnline struct{}

func (*defaultBuilderOnline) Build() Configuration {
	logger.Level = zap.DebugLevel
	logger.Target = logger.Console

	return Configuration{
		//MySQL:                "root:@tcp(127.0.0.1:3306)/im?charset=utf8&parseTime=true",
		MySQL:                "root:adminRoot@8888SecretPwd@tcp(127.0.0.1:3307)/im?charset=utf8&parseTime=true",
		RedisHost:            "127.0.0.1:6379",
		RedisPassword:        "",
		PushRoomSubscribeNum: 100,
		PushAllSubscribeNum:  100,

		ConnectLocalAddr:     "127.0.0.1:8000",
		ConnectRPCListenAddr: ":8000",
		ConnectTCPListenAddr: ":8001",
		ConnectWSListenAddr:  ":8002",

		LogicRPCListenAddr:    ":8010",
		BusinessRPCListenAddr: ":8020",
		FileHTTPListenAddr:    "8030",

		ConnectIntClientBuilder: func() pb.ConnectIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///127.0.0.1:8000", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, picker.AddrPickerName)))
			if err != nil {
				panic(err)
			}
			return pb.NewConnectIntClient(conn)
		},
		LogicIntClientBuilder: func() pb.LogicIntClient {
			//addrs:///docker.for.mac.host.internal:8010
			conn, err := grpc.DialContext(context.TODO(), "addrs:///127.0.0.1:8010", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewLogicIntClient(conn)
		},
		BusinessIntClientBuilder: func() pb.BusinessIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///127.0.0.1:8020", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewBusinessIntClient(conn)
		},
		//@ms:
		BusinessExtClientBuilder: func() pb.BusinessExtClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///127.0.0.1:8020", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewBusinessExtClient(conn)
		},
		LogicExtClientBuilder: func() pb.LogicExtClient {
			//addrs:///docker.for.mac.host.internal:8010
			conn, err := grpc.DialContext(context.TODO(), "addrs:///127.0.0.1:8010", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewLogicExtClient(conn)
		},
	}
}
