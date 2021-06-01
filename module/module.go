package module

import (
	"content-grpc/module/test"
	"google.golang.org/grpc"
)

func RegisterModule(server *grpc.Server) {
	test.Init(server)
}

