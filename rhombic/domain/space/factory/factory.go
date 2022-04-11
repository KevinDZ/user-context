package factory

import (
	"user-context/rhombic/domain"
	"user-context/rhombic/domain/space/service"
	"user-context/rhombic/domain/space/vo"
)

// Factory 继承父类：聚合根隐性依赖
type Factory struct {
	Root        *domain.AggregateRoot
	Application *vo.ValueObject
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
	if factory.Application != nil {
		return
	}
	factory.Client = service.NewSpaceService() // 实例化空间的领域服务
	if factory.Client == nil {
		return
	}
	factory.Application = &vo.ValueObject{}
	return true
}
