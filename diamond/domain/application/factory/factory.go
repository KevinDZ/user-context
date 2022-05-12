package factory

import (
	"strings"
	"user-context/diamond/domain"
	"user-context/diamond/domain/application/service"
	"user-context/diamond/domain/application/vo"
)

// Factory 继承父类：聚合根隐性依赖
type Factory struct {
	Aggregate   domain.AggregateRoot
	Application vo.Application
	Service     *service.Service
}

// InstanceApplicationAggregate 实例化聚合
func InstanceApplicationAggregate(rootID string) *Factory {
	return &Factory{Aggregate: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if strings.Replace(factory.Aggregate.GetAggregateRootID(), " ", "", -1) == "" {
		return
	}
	factory.Application = vo.Application{}
	return true
}
