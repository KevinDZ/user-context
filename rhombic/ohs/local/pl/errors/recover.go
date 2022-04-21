package errors

import (
	"fmt"
	"user-context/utils/common"
)

func Recover(types, method, clientType, sitCode, ip, proxies string, code int64) {
	if r := recover(); r != nil {
		switch types {
		case "system":
			fmt.Println(r)
		default:
			common.Panicln(method, clientType, sitCode, ip, proxies, fmt.Sprintf("%v", r), code)
		}
	}
}
