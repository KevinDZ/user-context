package factory

import (
	"user-context/diamond/domain"
	"user-context/diamond/domain/organization/service"
	"user-context/diamond/domain/organization/vo"
)

// Factory 继承AggregateRoot父类：聚合根隐性依赖
type Factory struct {
	Aggregate    *domain.AggregateRoot
	Organization *vo.Organization
	Client       *service.Service
}

// InstanceOrganizationAggregate 实例化聚合
func InstanceOrganizationAggregate(rootID string) *Factory {
	return &Factory{Aggregate: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Aggregate.RootID) == 0 {
		return
	}
	factory.Organization = &vo.Organization{}
	return true
}
