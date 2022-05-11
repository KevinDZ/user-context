package factory

import (
	"user-context/rhombic/domain"
	"user-context/rhombic/domain/application/service"
	"user-context/rhombic/domain/application/vo"
)

// Factory 继承父类：聚合根隐性依赖
type Factory struct {
	Root        *domain.AggregateRoot
	Application *vo.Application
	Client      *service.Service
}

// InstanceApplicationAggregate 实例化聚合
func InstanceApplicationAggregate(rootID string) *Factory {
	return &Factory{Root: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Root.RootID) == 0 {
		return
	}
	factory.Application = &vo.Application{}
	return true
}
