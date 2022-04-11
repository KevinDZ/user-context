package repositories

import (
	"user-context/rhombic/ohs/local/pl/vo"
)

func IsExist(request vo.LoginRequest) (ok bool) {
	// 1.通过 key 判断是否存在
	// 2.通过 value 判断是否正确
	return
}

func Delete(request vo.LogoutRequest) (err error) {
	// 1.通过 key 删除
	return
}
