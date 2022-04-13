package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGRPC() {
	//启动grpc 服务端服务
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//space := clients.NewSpaceAdapter()
	//dto.RegisterSpaceServiceServer(grpcServer, &space)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
