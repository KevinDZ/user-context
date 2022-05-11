package services

import (
	"errors"
	"user-context/diamond/acl/adapters/clients"
	"user-context/diamond/acl/adapters/pl/dao"
	"user-context/diamond/acl/adapters/publishers"
	"user-context/diamond/acl/adapters/repositories"
	"user-context/diamond/domain/account/factory"
	"user-context/diamond/ohs/local/pl/vo"
)

// LoginAppService 登录应用服务
func LoginAppService(request vo.LoginRequest) (user dao.UserDAO, token string, err error) {
	// 实例化数据库
	db := repositories.NewDAO("db")
	// 1. 用户账户表:判断账户是否存在
	// 	1.1 CQRS原则 - 读操作：直接获取，不经过领域模型
	if err = db.AccountDAO(request.UserID); err != nil {
		// 不存在,返回
		return
	}
	// 2.写账户登录表 - MySQL
	//	2.1 写操作：无领域模型
	if err = db.LoginRecord(request); err != nil {
		return
	}
	token, err = db.CreateToken()
	// 3.发送redis保存
	if err != nil {
		return
	}
	return
}

// LogoutAppService 登出应用服务
func LogoutAppService(request vo.LogoutRequest) (err error) {
	// 清除redis保存的account token
	db := repositories.NewDAO("redis")
	defer db.Redis.Close()
	err = db.Redis.Del(request.ID).Err()
	if err != nil {
		return
	}
	return
}

// RegisteredAppService 注册账户应用服务
func RegisteredAppService(request vo.RegisteredRequest) (result vo.LoginRequest, err error) {
	// 0.通过聚合根ID，实例化聚合根
	account := factory.InstanceAccountAggregate("") // rootID = entity ID
	// 1.实例化聚合和领域服务
	// 1.1 实例化领域服务：端口与适配器实现
	account.Service.Repository = repositories.NewAccountAdapter()
	if account.Service.Repository == nil {
		err = errors.New("repository instance failed")
		return
	}
	account.Service.Client = clients.NewUUIDAdapter()
	if account.Service.Client == nil {
		err = errors.New("client instance failed")
		return
	}

	// 1.2 实例化聚合
	if err = account.InstanceOf(); err != nil {
		return
	}
	// 2.填充聚合内可选参数
	account.WithAccountOptions(request.NickName, request.PassWord, request.Mobile, request.Email, request.BandID)

	// 3.考虑手机验证码校验 直接调用 不进领域层
	db := repositories.NewDAO("redis")
	defer db.Redis.Close() // 关闭redis连接
	if err = db.MobileVerify(request.Mobile, request.MobileCaptcha); err != nil {
		return
	}
	// 4.调用注册的领域服务
	if err = account.Registered(); err != nil {
		return
	}

	// 应用事件必须同步实现，不能异步通知
	// 5.1 事件实例化
	name := "" // TODO what is name ?
	account.Service.Publisher = publishers.NewAccountEvent(name)
	if account.Service.Publisher == nil {
		err = errors.New("publisher instance failed")
		return
	}
	// 5.2 事件关闭连接
	defer account.Service.Publisher.Close()        // 后关闭 connection
	defer account.Service.Publisher.ChannelClose() // 先关闭 channel

	// 6.发布注册应用事件：空间、套餐
	if err = account.RegisteredEvent(); err != nil {
		return // 注册事件失败， 返回注册失败
	}
	result = vo.LoginRequest{}
	return
}

// 事件风暴梳理出来的都可以作为功能
// 修改个人信息
// 校验更新密码
// 校验用户账户密码
// 用户绑定信息修改
// 个人信息更新事件
