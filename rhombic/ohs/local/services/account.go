package services

import (
	"github.com/spf13/viper"
	"user-context/rhombic/acl/adapters/repositories"
	"user-context/rhombic/domain/account/factory"
	"user-context/rhombic/ohs/local/pl/vo"
)

// LoginAppService 登录应用服务
func LoginAppService(request vo.LoginRequest) (result vo.LoginRespond, err error) {
	dao := repositories.NewDAO()
	// 1.判断账户是否存在 - redis获取用户信息
	// 	1.1 CQRS原则 - 读操作：直接获取，不经过领域模型
	if ok := dao.IsExist(request); !ok {
		// 不存在,返回
		return
	}
	// 2.写账户登录表 - MySQL
	//	2.1 写操作：无领域模型
	if err = dao.LoginRecord(request); err != nil {
		// err报错,返回
		return
	}
	// 3.读用户账户表 - MySQL
	//	3.1 CQRS原则 - 读操作：直接通过端口获取，不经过领域模型
	result = dao.AccountDAO(request)
	return
}

// LogoutAppService 登出应用服务
func LogoutAppService(request vo.LogoutRequest) (err error) {
	dao := repositories.NewDAO()
	// 清除redis保存的account token
	err = dao.Delete(request)
	if err != nil {
		return
	}
	return
}

// RegisteredAppService 注册账户应用服务
func RegisteredAppService(request vo.RegisteredRequest) {
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
	// 4.发布注册应用事件
	if account.Service.Publisher.Registered(viper.GetString("channel.user"), account.Entity.Event) {
		// 注册事件失败， 返回注册失败
		return
	}
}

// 事件风暴梳理出来的都可以作为功能
// 修改个人信息
// 校验更新密码
// 校验用户账户密码
// 用户绑定信息修改
// 个人信息更新事件
