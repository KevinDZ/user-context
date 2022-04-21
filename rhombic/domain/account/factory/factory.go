package factory

import (
	"errors"
	"github.com/spf13/viper"
	"time"
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
	factory.Entity.NickName = name
	factory.Entity.PassWord = passwd
	factory.Entity.Mobile = phone
	factory.Entity.Email = email
	factory.Entity.ThirdPartyID = thirdParty
	return factory
}

// InstanceOf 实例化聚合和启动领域服务
func (factory *Factory) InstanceOf() (err error) {
	if len(factory.Root.RootID) == 0 {
		err = errors.New("aggregate root id is null")
		return
	}
	if factory.Entity != nil {
		err = errors.New("entity is exist")
		return
	}
	factory.Service = service.NewAccountService()
	if factory.Service == nil || factory.Service.Repository == nil || factory.Service.Publisher == nil {
		err = errors.New("service not started")
		return
	}
	factory.Entity = &entity.Entity{ID: factory.Root.GetAggregateRootID()}
	return
}

// InstanceOfEvent 实例化事件
func (factory *Factory) InstanceOfEvent() (err error) {
	factory.Service = service.NewAccountEvent()
	if factory.Service.Publisher == nil {
		err = errors.New("publisher instance failed")
	}
	if factory.Service.Repository == nil {
		err = errors.New("repository instance failed")
	}
	return
}

// Registered 调用注册的领域服务
func (factory *Factory) Registered() (err error) {
	// 用户唯一身份标识：获取uuid
	factory.Entity.ID = factory.Service.Client.GetUUID()
	// 领域服务：用户注册的持久化
	if err = factory.Service.Registered(*factory.Entity); err != nil {
		err = errors.New("domain service save persistence data failed")
		return
	}
	// 设置实体领域行为：注册事件
	// 用户初始化配置分配：空间、套餐
	// 1.发布通道
	factory.Channel = viper.GetString("channel.user")
	// 2.事件消息
	factory.Event = map[string]string{factory.Entity.ID: viper.GetString("event.registered")}
	return

}

// RegisteredEvent 注册事件，失败后重试机制
func (factory *Factory) RegisteredEvent(count int64) (err error) {
RETRY:
	err = factory.Service.Publisher.Registered(factory.Channel, factory.Event)
	if err != nil {
		for count < 5 {
			time.Sleep(time.Duration(2<<count) * time.Second)
			count++
			goto RETRY
		}
	}
	return
}
