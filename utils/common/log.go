package common

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger 服务日志服务，可适配其他日志组件
type Logger interface {
	Infoln(method, clientType, sitCode, client, service, proxies, message string)
	Debugln(method, clientType, sitCode, client, service, proxies, message string)
	Errorln(method, clientType, sitCode, client, service, proxies, message string, code int64)
	Warnln(method, clientType, sitCode, client, service, proxies, message string)
	Traceln(method, clientType, sitCode, client, service, proxies, message string)
	Panicln(method, clientType, sitCode, client, service, proxies, message string, code int64)
	Fatalln(method, clientType, sitCode, client, service, proxies, message string, code int64)
}

var (
	logOnce sync.Once
	l       Logger
	logDir  string
	logName string
)

// NewLogger 获取日志句柄
func NewLogger() Logger {
	initLogger()

	return l
}

type serverLog struct {
	logger *logrus.Logger
}

// Infoln 普通信息
func (l *serverLog) Infoln(method, clientType, sitCode, client, service, proxies, message string) {
	if l.logger == nil {
		return
	}
	format := Format{
		Level:      "Infoln",
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     client,
		Service:    service,
		Proxies:    proxies,
		Message:    message,
	}
	l.logger.Infof("%+v\n", format)
}

// Warnln 警告信息
func (l *serverLog) Warnln(method, clientType, sitCode, client, service, proxies, message string) {
	if l.logger == nil {
		return
	}
	format := Format{
		Level:      "Warnln",
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     client,
		Service:    service,
		Proxies:    proxies,
		Message:    message,
	}
	l.logger.Warnf("%+v\n", format)
}

// Errorln 错误信息
func (l *serverLog) Errorln(method, clientType, sitCode, client, service, proxies, message string, code int64) {
	if l.logger == nil {
		return
	}
	format := Format{
		Level:      "Errorln",
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     client,
		Service:    service,
		Proxies:    proxies,
		Message:    message,
		Code:       code,
	}
	l.logger.Errorf("%+v\n", format)
}

// Debugln 调试信息
func (l *serverLog) Debugln(method, clientType, sitCode, client, service, proxies, message string) {
	if l.logger == nil {
		return
	}
	format := Format{
		Level:      "Debugf",
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     client,
		Service:    service,
		Proxies:    proxies,
		Message:    message,
	}
	l.logger.Debugf("%+v\n", format)
}

// Traceln 跟踪信息
func (l *serverLog) Traceln(method, clientType, sitCode, client, service, proxies, message string) {
	if l.logger == nil {
		return
	}
	format := Format{
		Level:      "Traceln",
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     client,
		Service:    service,
		Proxies:    proxies,
		Message:    message,
	}
	l.logger.Tracef("%+v\n", format)
}

// Fatalln 致命错误
func (l *serverLog) Fatalln(method, clientType, sitCode, client, service, proxies, message string, code int64) {
	if l.logger == nil {
		return
	}
	format := Format{
		Level:      "Fatalln",
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     client,
		Service:    service,
		Proxies:    proxies,
		Message:    message,
		Code:       code,
	}
	l.logger.Fatalf("%+v\n", format)
}

// Panicln 恐慌错误
func (l *serverLog) Panicln(method, clientType, sitCode, client, service, proxies, message string, code int64) {
	if l.logger == nil {
		return
	}
	format := Format{
		Level:      "Panicln",
		Now:        time.Now().Unix(),
		Method:     method,
		ClientType: clientType,
		SiteCode:   sitCode,
		Client:     client,
		Service:    service,
		Proxies:    proxies,
		Message:    message,
		Code:       code,
	}
	l.logger.Panicf("%+v\n", format)
}

func initLogger() {
	logOnce.Do(func() {
		// config := NewConfig()
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		logHandle := &serverLog{}
		logHandle.logger = logrus.New()
		logHandle.logger.SetFormatter(&logrus.JSONFormatter{})
		logout := os.Getenv("LOGOUT")
		if len(logout) > 0 {
			logHandle.logger.SetOutput(os.Stdout)
		} else {
			// logDir = config.LogDir
			// logName = config.LogName
			// if logDir == "" {
			// 	logDir = GetProjectAbPathByCaller()
			// } else {
			// 	err := os.MkdirAll(logDir, 0750)
			// 	if err != nil {
			// 		fmt.Println("mkdir err", err)
			// 	}
			// }
			// if logName == "" {
			// 	logName = "server.log"
			// }
			// logFileName := path.Join(logDir, logName)
			// logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
			// if err != nil {
			// 	fmt.Println("open log file err", err)
			// }
			// logHandle.logger.SetOutput(logFile)
			// l = logHandle
		}
	})
}
