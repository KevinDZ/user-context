package server

import (
	"user_context/utils/gin/cmd"
)

// StartGin gin 框架 http 服务启动
func StartGin() {
	cmd.Execute()
}