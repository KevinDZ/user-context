package services

import (
	"user-context/rhombic/acl/adapters/repositories"
	"user-context/rhombic/domain/account/factory"
	"user-context/rhombic/ohs/local/pl"
)

// LoginAppService 登录应用服务
func LoginAppService(request pl.LoginRequest) (result pl.LoginRespond, err error) {
	// 1.判断账户是否存在 - redis获取用户信息
	// 	1.1 CQRS原则 - 读操作：直接获取，不经过领域模型
	if ok := repositories.IsExist(request); !ok {
		// 不存在,返回
		return
	}
	// 2.写账户登录表 - MySQL
	//	2.1 写操作：无领域模型
	if err = repositories.LoginRecord(request); err != nil {
		// err报错,返回
		return
	}
	// 3.读用户账户表 - MySQL
	//	3.1 CQRS原则 - 读操作：直接通过端口获取，不经过领域模型
	result = repositories.AccountDAO(request)
	return
}

// LogoutAppService 登出应用服务
func LogoutAppService(request pl.LogoutRequest) (err error) {
	// 清除redis保存的account token
	err = repositories.Delete(request)
	if err != nil {
		return
	}
	return
}

// RegisteredAppService 注册账户应用服务
func RegisteredAppService(request pl.RegisteredRequest) {
	rootID := ""
	// 0.通过聚合根ID，实例化聚合根
	account := factory.InstanceAccountAggregate(rootID)
	// 1.填充聚合内可选参数
	account.WithAccountOptions(request.Name, request.PassWord, request.Phone, request.Email, "")
	// 2.通过聚合，实例化聚合和领域服务
	if account.InstanceOf() {
		return
	}
	// 3.调用注册的领域服务
	if account.Registered() {
		return
	}
}
