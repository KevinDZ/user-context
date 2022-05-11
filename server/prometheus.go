package server

import (
	"fmt"
	"net/http"
	"os"
	"user-context/diamond/ohs/local/pl/errors"
)

// startPrometheus 监控 服务健康
func startPrometheus() {
	defer errors.Recover("system", "", "", "", "", "", 500)
	fmt.Println("prometheus start...")
	http.HandleFunc("/health", health)
	http.HandleFunc("/dependency", serviceDependency)
	http.ListenAndServe("127.0.0.1:8081", nil)
}

// 监控服务健康
func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("health"))
	return
}

// 服务依赖健康
func serviceDependency(w http.ResponseWriter, r *http.Request) {
	var err error
	if err != nil {
		// 依赖断开，直接服务关闭
		os.Exit(1)
	}
}
