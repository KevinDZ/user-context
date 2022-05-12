package factory

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
	"user-context/diamond/domain"
	"user-context/diamond/domain/account/entity"
	"user-context/diamond/domain/account/service"
)

type Factory struct {
	// 聚合根
	Aggregate domain.AggregateRoot
	// 聚合
	Account entity.Account
	// 领域服务
	Service *service.Service

	// 应用事件
	Event map[string]string
}

// InstanceAccountAggregate 实例化聚合
func InstanceAccountAggregate(rootID string) *Factory {
	return &Factory{Aggregate: domain.NewAggregateRoot(rootID)}
}

// WithAccountOptions init account aggregate with account params
func (factory *Factory) WithAccountOptions(name, passwd, phone, email, thirdParty string) *Factory {
	factory.Account.SetID(factory.Aggregate.GetAggregateRootID())
	factory.Account.SetNickName(name)
	factory.Account.SetPassWord(passwd)
	factory.Account.SetMobile(phone)
	factory.Account.SetEmail(email)
	factory.Account.SetThirdPartyID(thirdParty)
	return factory
}

// InstanceOf 实例化聚合和启动领域服务
func (factory *Factory) InstanceOf() (err error) {
	if strings.Replace(factory.Aggregate.GetAggregateRootID(), " ", "", -1) == "" {
		// 用户唯一身份标识：获取uuid
		factory.Aggregate.SetAggregateRootID(factory.Service.GetUUID())
	}
	factory.Account = entity.Account{}
	return
}

// Registered 调用注册的领域服务
func (factory *Factory) Registered() (err error) {
	// 领域服务：用户注册的持久化
	if err = factory.Service.Registered(factory.Account); err != nil {
		err = errors.New("domain service save persistence data failed")
		return
	}
	// 设置实体领域行为：注册事件
	// 用户初始化配置分配：空间、套餐
	// 1.事件消息
	factory.Event = map[string]string{factory.Account.GetID(): viper.GetString("event.registered")}
	return

}
