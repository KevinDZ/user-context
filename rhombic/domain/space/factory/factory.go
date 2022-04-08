package factory

import (
	"user_context/rhombic/domain"
	"user_context/rhombic/domain/space/service"
	"user_context/rhombic/domain/space/vo"
)


// Factory 继承父类：聚合根隐性依赖
type Factory struct {
	Root *domain.AggregateRoot
	Application *vo.ValueObject
	Client *service.Service
}

// InstanceSpaceAggregate 实例化聚合
func InstanceSpaceAggregate(rootID string) *Factory {
	return &Factory{Root: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Root.RootID) == 0 {return}
	if factory.Application != nil {return}
	factory.Client = service.NewSpaceService()
	if factory.Client == nil {return}
	factory.Application = &vo.ValueObject{}
	return true
}