package core

import (
	"content-grpc/global"
	"content-grpc/module"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)


func InitGrpcServer(port string) (server *grpc.Server) {
	server = grpc.NewServer()

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		global.LOG.Error("net.Listen", zap.Any("err", err))
	}

	//注册路由
	module.RegisterModule(server)

	fmt.Printf("grpc服务已启动 端口号：%s", port)

	server.Serve(lis)

	return server
}



