package main

import (
	"flag"
	"fmt"

	"imooc.com/traning/user/rpc/internal/config"
	"imooc.com/traning/user/rpc/internal/server"
	"imooc.com/traning/user/rpc/internal/svc"
	"imooc.com/traning/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// s.AddUnaryInterceptors()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
