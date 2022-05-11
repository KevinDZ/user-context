package factory

import (
	"user-context/diamond/domain"
	"user-context/diamond/domain/space/service"
	"user-context/diamond/domain/space/vo"
)

// Factory 继承父类：聚合根隐性依赖
type Factory struct {
	Root        *domain.AggregateRoot
	Application *vo.Space
	Client      *service.Service
}

// InstanceSpaceAggregate 实例化聚合
func InstanceSpaceAggregate(rootID string) *Factory {
	return &Factory{Root: domain.NewAggregateRoot(rootID)}
}

// InstanceOf 实例化聚合和领域服务
func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Root.RootID) == 0 {
		return
	}
	factory.Application = &vo.Space{}
	return true
}
