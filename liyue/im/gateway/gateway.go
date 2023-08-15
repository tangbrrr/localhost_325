package main

import (
	"flag"
	"fmt"

	"github.com/tangbo/twatt/liyue/im/gateway/internal/config"
	"github.com/tangbo/twatt/liyue/im/gateway/internal/server"
	"github.com/tangbo/twatt/liyue/im/gateway/internal/svc"
	"github.com/tangbo/twatt/liyue/im/gateway/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	gatewaySvc := server.NewGatewayServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterGatewayServer(grpcServer, gatewaySvc)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// TODO: 启动长链接服务
	gatewaySvc.Listen()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
