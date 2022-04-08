package service

//领域服务：主要的业务规则编写
import (
	pubAdapter "user-context/rhombic/acl/adapters/publishers"
	repoAdapter "user-context/rhombic/acl/adapters/repositories"
	"user-context/rhombic/acl/ports/publishers"
	"user-context/rhombic/acl/ports/repositories"
	"user-context/rhombic/domain/account/entity"
)

type Service struct {
	Repository repositories.AccountRepository //资源库端口
	Publisher  publishers.AccountPublisher    //发布者端口
}

func NewAccountService() *Service {
	channel := ""
	return &Service{
		repoAdapter.NewAccountAdapter(),
		pubAdapter.NewAccountEvent(channel),
	}
}

// Registered 注册账户
func (service *Service) Registered(entity entity.Entity) (ok bool) {
	if ok = service.Repository.CheckIsExist(entity); ok {
		return !ok
	} // 账户存在，返回注册失败
	if err := service.Repository.Insert(entity); err != nil {
		return
	} // 账户写入失败，返回注册失败
	if ok = service.Publisher.Registered(entity.ID, entity.Event); !ok {
		return ok
	} // 注册事件失败， 返回注册失败
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
