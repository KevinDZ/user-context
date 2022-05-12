package factory

import (
	"strings"
	"user-context/diamond/domain"
	"user-context/diamond/domain/package/service"
	"user-context/diamond/domain/package/vo"
)

// Factory 继承AggregateRoot父类：聚合根隐性依赖
type Factory struct {
	Aggregate domain.AggregateRoot
	Package   vo.Package
	Service   *service.Service
}

// InstancePackageAggregate 实例化聚合
func InstancePackageAggregate(rootID string) *Factory {
	return &Factory{Aggregate: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if strings.Replace(factory.Aggregate.GetAggregateRootID(), " ", "", -1) == "" {
		return
	}
	factory.Package = vo.Package{}
	return true
}
