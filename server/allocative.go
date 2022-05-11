package server

import (
	"fmt"
	"os"
	"time"
	"user-context/diamond/ohs/local/pl/errors"
)

func StartServe() {
	// 配置中心获取
	go getAllocateCenter()
	// prometheus 监控服务启动
	go startPrometheus()
}

// GetAllocateCenter 配置中心服务
func getAllocateCenter() {
	defer errors.Recover("system", "", "", "", "", "", 500)
	var err error
	// 获取配置中心
	for {
		time.Sleep(time.Duration(10) * time.Second)
		// TODO 配置文件已用完，重新读取：热加载
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

	}
}
