package services

import (
	"fmt"
	"user-context/rhombic/acl/adapters/pl/dao"
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
func RegisteredAppService(request vo.RegisteredRequest) (result vo.LoginRequest, err error) {
	// 0.通过聚合根ID，实例化聚合根
	account := factory.InstanceAccountAggregate(request.RootID)
	// 1.通过聚合，实例化聚合和领域服务
	if err = account.InstanceOf(); err != nil {
		fmt.Println("account.InstanceOf() error: ", err)
		return
	}
	// 2.填充聚合内可选参数
	account.WithAccountOptions(request.NickName, request.PassWord, request.Mobile, request.Email, "")
	// 3.考虑手机验证码校验 直接调用 不进领域层
	if err = dao.NewRedisDAO().MobileVerify(request.Mobile, request.MobileCaptcha); err != nil {
		fmt.Println("account.MobileVerify() error: ", err)
		return
	}
	// 4.调用注册的领域服务
	if err = account.Registered(); err != nil {
		fmt.Println("account.Registered() error: ", err)
		return
	}
	// 5.事件实例化
	if err = account.InstanceOfEvent(); err != nil {
		fmt.Println("account.InstanceOfEvent() error: ", err)
		return
	}
	// 6.发布注册应用事件
	if err = account.RegisteredEvent(); err != nil {
		fmt.Println("account.RegisteredEvent() error: ", err)
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
