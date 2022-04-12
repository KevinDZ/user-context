package factory

import (
	"github.com/spf13/viper"
	"user-context/rhombic/domain"
	"user-context/rhombic/domain/account/entity"
	"user-context/rhombic/domain/account/service"
)

type Factory struct {
	// 聚合根
	Root *domain.AggregateRoot
	// 聚合
	Entity *entity.Entity
	// 领域服务
	Service *service.Service

	// 应用事件
	Channel string
	Event   map[string]string
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
	// 用户唯一身份标识：获取uuid
	factory.Entity.ID = factory.Service.Client.GetUUID()
	// 领域服务：用户注册的持久化
	if ok = factory.Service.Registered(*factory.Entity); !ok {
		return
	}
	// 设置实体领域行为：注册事件
	// 1.发布通道
	factory.Channel = viper.GetString("channel.user")
	// 2.事件消息
	factory.Event = map[string]string{factory.Entity.ID: viper.GetString("event.registered")}
	return
}
