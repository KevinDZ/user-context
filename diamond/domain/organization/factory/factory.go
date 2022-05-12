package factory

import (
	"strings"
	"user-context/diamond/domain"
	"user-context/diamond/domain/organization/service"
	"user-context/diamond/domain/organization/vo"
)

// Factory 继承AggregateRoot父类：聚合根隐性依赖
type Factory struct {
	Aggregate    domain.AggregateRoot
	Organization vo.Organization
	Service      *service.Service
}

// InstanceOrganizationAggregate 实例化聚合
func InstanceOrganizationAggregate(rootID string) *Factory {
	return &Factory{Aggregate: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if strings.Replace(factory.Aggregate.GetAggregateRootID(), " ", "", -1) == "" {
		return
	}
	factory.Organization = vo.Organization{}
	return true
}
