package service

//领域服务：主要的业务规则编写
import (
	"user-context/diamond/acl/ports/clients"
	"user-context/diamond/acl/ports/repositories"
	"user-context/diamond/domain/account/entity"
)

type Service struct {
	// 资源服务
	UUID    clients.UUIDClient             //客户端端口
	Account repositories.AccountRepository //资源库端口
}

// GetUUID 获取UUID
func (service *Service) GetUUID() string {
	return service.UUID.GetUUID()
}

// Registered 注册账户 -- 业务规则
func (service *Service) Registered(entity entity.Account) (err error) {
	if err = service.Account.CheckIsExist(entity); err != nil {
		return
	} // 账户存在，返回注册失败
	if err = service.Account.Insert(entity); err != nil {
		return
	} // 账户写入失败，返回注册失败
	return
}

// Query 查询账户
func (service *Service) Query(id string) *entity.Account {
	return service.Account.Query(id)
}

// Verify 验证账户
func (service *Service) Verify(entity entity.Account) error {
	return service.Account.CheckIsExist(entity)
}

// 事件风暴梳理出来的都可以作为功能

// ModifyPersonInformation 修改个人信息
func (service *Service) ModifyPersonInformation(entity entity.Account) error {
	return service.Account.Update(entity)
}

// 校验更新密码
// 校验用户账户密码
// 用户绑定信息修改
// 个人信息更新事件
