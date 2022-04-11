package repositories

import (
	"user-context/rhombic/acl/adapters/pl/dao"
	"user-context/rhombic/ohs/local/pl/vo"
)

//LoginRecord 用户登录表 - 写操作
func LoginRecord(request vo.LoginRequest) (err error) {
	var dao dao.LoginRecordDAO
	_ = dao
	return
}

// AccountDAO CQRS原则：账户数据访问对象  - 读操作
func AccountDAO(request vo.LoginRequest) (result vo.LoginRespond) {
	// 用户数据表 - 读操作
	dao := dao.PostgreSQLDAO{ID: request.ID}
	_ = dao
	return
}
