package factory

import (
	"user_context/rhombic/domain"
	"user_context/rhombic/domain/application/service"
	"user_context/rhombic/domain/application/vo"
)

// Factory 继承父类：聚合根隐性依赖
type Factory struct {
	Root *domain.AggregateRoot
	Application *vo.ValueObject
	Client *service.Service
}

// InstanceApplicationAggregate 实例化聚合
func InstanceApplicationAggregate(rootID string) *Factory {
	return &Factory{Root: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Root.RootID) == 0 {return}
	if factory.Application != nil {return}
	factory.Client = service.NewApplicationService()
	if factory.Client == nil {return}
	factory.Application = &vo.ValueObject{}
	return true
}