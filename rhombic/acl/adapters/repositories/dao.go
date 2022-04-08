package repositories

import (
	pl2 "user-context/rhombic/ohs/local/pl"
)

//LoginRecord 用户登录表 - 写操作
func LoginRecord(request pl2.LoginRequest) (err error) {
	var dao pl2.LoginRecordDAO
	_ = dao
	return
}

// AccountDAO CQRS原则：账户数据访问对象  - 读操作
func AccountDAO(request pl2.LoginRequest) (result pl2.LoginRespond) {
	// 用户数据表 - 读操作
	dao := pl2.MysqlDAO{ID: request.ID}
	_ = dao
	return
}
