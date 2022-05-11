package factory

import (
	"user-context/diamond/domain"
	"user-context/diamond/domain/package/service"
	"user-context/diamond/domain/package/vo"
)

// Factory 继承AggregateRoot父类：聚合根隐性依赖
type Factory struct {
	Aggregate *domain.AggregateRoot
	Package   *vo.Package
	Client    *service.Service
}

// InstancePackageAggregate 实例化聚合
func InstancePackageAggregate(rootID string) *Factory {
	return &Factory{Aggregate: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Aggregate.RootID) == 0 {
		return
	}
	factory.Package = &vo.Package{}
	return true
}
