package services

import (
	"user-context/rhombic/acl/adapters/pl/dao"
	"user-context/rhombic/acl/adapters/repositories"
	"user-context/rhombic/domain/account/factory"
	"user-context/rhombic/ohs/local/pl/vo"
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
	account := factory.InstanceAccountAggregate(request.RootID)
	// 1.通过聚合，实例化聚合和领域服务
	if err = account.InstanceOf(); err != nil {
		return
	}
	// 2.填充聚合内可选参数
	account.WithAccountOptions(request.NickName, request.PassWord, request.Mobile, request.Email, request.BandID)
	// 3.考虑手机验证码校验 直接调用 不进领域层
	db := repositories.NewDAO("redis")
	defer db.Redis.Close()
	if err = db.MobileVerify(request.Mobile, request.MobileCaptcha); err != nil {
		return
	}
	// 4.调用注册的领域服务
	if err = account.Registered(); err != nil {
		return
	}

	// 保证应用事件消息不丢失
	// 5.事件实例化
	if err = account.InstanceOfEvent(); err != nil {
		return
	}
	// 6.记录注册事件，保证事件消息不丢失
	db = repositories.NewDAO("db")
	if err = db.EventCreate(); err != nil {
		return
	}
	// 7.发布注册应用事件：空间、套餐
	if err = account.RegisteredEvent(); err != nil {
		return // 注册事件失败， 返回注册失败
	}
	// 8.注册应用事件完成
	if err = db.EventComplete(); err != nil {
		return
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
