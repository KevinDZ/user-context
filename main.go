package main

import (
	"user_context/server"
)

func main() {
	// grpc 服务启动
	go server.StartGRPC()
	// gin 框架 http 服务启动
	server.StartGin()
}
