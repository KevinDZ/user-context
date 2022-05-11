package factory

import (
	"user-context/rhombic/domain"
	"user-context/rhombic/domain/organization/service"
	"user-context/rhombic/domain/organization/vo"
)

// Factory 继承AggregateRoot父类：聚合根隐性依赖
type Factory struct {
	Root         *domain.AggregateRoot
	Organization *vo.Organization
	Client       *service.Service
}

// InstanceOrganizationAggregate 实例化聚合
func InstanceOrganizationAggregate(rootID string) *Factory {
	return &Factory{Root: domain.NewAggregateRoot(rootID)}
}

func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Root.RootID) == 0 {
		return
	}
	factory.Organization = &vo.Organization{}
	return true
}
