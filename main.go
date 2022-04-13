package main

import (
	"user-context/server"
)

func main() {
	// grpc 服务启动
	//go server.StartGRPC() // TODO 用户上下文  不需要grpc服务端被调用
	// gin 框架 http 服务启动
	server.StartGin()
}
