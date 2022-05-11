package main

import (
	"runtime"
	"user-context/server"
)

func main() {
	// 开启多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 依赖服务
	go server.StartServe()

	// grpc 服务启动
	//go server.StartGRPC() // TODO 用户上下文  不需要grpc服务端被调用

	// gin 框架 http 服务启动
	server.StartGin()
}
