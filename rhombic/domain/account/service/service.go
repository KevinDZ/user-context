package service

//领域服务：主要的业务规则编写
import (
	cliAdapter "user-context/rhombic/acl/adapters/clients"
	pubAdapter "user-context/rhombic/acl/adapters/publishers"
	repoAdapter "user-context/rhombic/acl/adapters/repositories"
	"user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/acl/ports/publishers"
	"user-context/rhombic/acl/ports/repositories"
	"user-context/rhombic/domain/account/entity"
)

type Service struct {
	Client     clients.UUIDClient             //客户端端口
	Repository repositories.AccountRepository //资源库端口
	Publisher  publishers.AccountPublisher    //发布者端口
}

func NewAccountService() *Service {
	return &Service{
		Client:     cliAdapter.NewUUIDAdapter(),
		Repository: repoAdapter.NewAccountAdapter(),
		Publisher:  pubAdapter.NewAccountEvent(),
	}
}

// Registered 注册账户 -- 业务规则
func (service *Service) Registered(entity entity.Entity) (ok bool) {
	if ok = service.Repository.CheckIsExist(entity); ok {
		return !ok
	} // 账户存在，返回注册失败
	if err := service.Repository.Insert(entity); err != nil {
		return
	} // 账户写入失败，返回注册失败
	return true
}

// Query 查询账户
func (service *Service) Query(id string) entity.Entity {
	return service.Repository.Query(id)
}

// Verify 验证账户
func (service *Service) Verify(entity entity.Entity) bool {
	return service.Repository.CheckIsExist(entity)
}

// 事件风暴梳理出来的都可以作为功能

// ModifyPersonInformation 修改个人信息
func (service *Service) ModifyPersonInformation(entity entity.Entity) error {
	return service.Repository.Update(entity)
}

// 校验更新密码
// 校验用户账户密码
// 用户绑定信息修改
// 个人信息更新事件
