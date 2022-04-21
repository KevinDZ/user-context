package common

import (
	"fmt"
	"time"
)

type Format struct {
	Level      string `json:"level" bson:"level"`
	Now        int64  `json:"now" bson:"now"`
	Method     string `json:"method" bson:"method"`
	Code       int64  `json:"code,omitempty" bson:"code"`
	ClientType string `json:"client_type" bson:"client_type"`
	SiteCode   string `json:"site_code" bson:"site_code"`
	Client     string `json:"client" bson:"client"`
	Service    string `json:"service" bson:"service"`
	Proxies    string `json:"proxies" bson:"proxies"`
	Message    string `json:"message" bson:"message"`
}

func com(level, method, clientType, sitCode, ip, proxies, message string) {
	format := Format{
		Level:      level,
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     ip,
		Service:    "user-context",
		Proxies:    proxies,
		Message:    message,
	}
	fmt.Printf("%+v\n", format)
}

func comCode(level, method, clientType, sitCode, ip, proxies, message string, code int64) {
	format := Format{
		Level:      level,
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     ip,
		Service:    "user-context",
		Proxies:    proxies,
		Message:    message,
		Code:       code,
	}
	fmt.Printf("%+v\n", format)
}

// Infoln 普通信息
func Infoln(method, clientType, sitCode, ip, proxies, message string) {
	com("Infoln", method, clientType, sitCode, ip, proxies, message)
}

// Warnln 警告信息
func Warnln(method, clientType, sitCode, ip, proxies, message string) {
	com("Warnln", method, clientType, sitCode, ip, proxies, message)
}

// Errorln 错误信息
func Errorln(method, clientType, sitCode, ip, proxies, message string, code int64) {
	comCode("Errorln", method, clientType, sitCode, ip, proxies, message, code)
}

// Debugln 调试信息
func Debugln(method, clientType, sitCode, ip, proxies, message string) {
	com("Debugln", method, clientType, sitCode, ip, proxies, message)
}

// Traceln 跟踪信息
func Traceln(method, clientType, sitCode, ip, proxies, message string) {
	com("Traceln", method, clientType, sitCode, ip, proxies, message)
}

// Fatalln 致命错误
func Fatalln(method, clientType, sitCode, ip, proxies, message string, code int64) {
	com("Fatalln", method, clientType, sitCode, ip, proxies, message)
}

// Panicln 恐慌错误
func Panicln(method, clientType, sitCode, ip, proxies, message string, code int64) {
	com("Panicln", method, clientType, sitCode, ip, proxies, message)
}
