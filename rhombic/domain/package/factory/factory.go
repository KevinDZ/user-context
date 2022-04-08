package factory

import (
	"user_context/rhombic/domain"
	"user_context/rhombic/domain/package/service"
	"user_context/rhombic/domain/package/vo"
)

// Factory 继承AggregateRoot父类：聚合根隐性依赖
type Factory struct {
	Root *domain.AggregateRoot
	Package *vo.ValueObject
	Client *service.Service
}

// InstancePackageAggregate 实例化聚合
func InstancePackageAggregate(rootID string) *Factory {
	return &Factory{Root: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Root.RootID) == 0 {return}
	if factory.Package != nil {return}
	factory.Client = service.NewPackageService()
	if factory.Client == nil {return}
	factory.Package = &vo.ValueObject{}
	return true
}