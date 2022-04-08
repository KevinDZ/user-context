package factory

import (
	"user-context/rhombic/domain"
	"user-context/rhombic/domain/account/entity"
	"user-context/rhombic/domain/account/service"
)

type Factory struct {
	Root    *domain.AggregateRoot
	Entity  *entity.Entity
	Service *service.Service
}

// InstanceAccountAggregate 实例化聚合
func InstanceAccountAggregate(rootID string) *Factory {
	return &Factory{Root: domain.NewAggregateRoot(rootID)}
}

// WithAccountOptions init account aggregate with account params
func (factory *Factory) WithAccountOptions(name, passwd, phone, email, thirdParty string) *Factory {
	factory.Entity.Name = name
	factory.Entity.PassWord = passwd
	factory.Entity.Phone = phone
	factory.Entity.Email = email
	factory.Entity.ThirdPartyID = thirdParty
	return factory
}

// InstanceOf 实例化聚合和启动领域服务
func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Root.RootID) == 0 {
		return
	}
	if factory.Entity != nil {
		return
	}
	factory.Service = service.NewAccountService()
	if factory.Service == nil || factory.Service.Repository == nil || factory.Service.Publisher == nil {
		return
	}
	factory.Entity = &entity.Entity{ID: factory.Root.GetAggregateRootID()}
	return true
}

// Registered 调用注册的领域服务
func (factory *Factory) Registered() (ok bool) {
	if factory.Service == nil || factory.Service.Repository == nil || factory.Service.Publisher == nil {
		return
	}
	ok = factory.Service.Registered(*factory.Entity)
	return
}
