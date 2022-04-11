package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGRPC() {
	//启动grpc服务
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
