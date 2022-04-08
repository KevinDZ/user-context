package repositories

import (
	"user_context/rhombic/ohs/local/pl"
)

func IsExist(request pl.LoginRequest) (ok bool) {
	// 1.通过 key 判断是否存在
	// 2.通过 value 判断是否正确
	return
}

func Delete(request pl.LogoutRequest) (err error) {
	// 1.通过 key 删除
	return
}