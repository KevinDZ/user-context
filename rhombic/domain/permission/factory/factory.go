package factory

import (
	"user-context/rhombic/domain"
	"user-context/rhombic/domain/permission/service"
	"user-context/rhombic/domain/permission/vo"
)

// Factory 继承父类：聚合根隐性依赖
type Factory struct {
	Root        *domain.AggregateRoot
	Application *vo.Permission
	Client      *service.Permission
}

// InstancePermissionAggregate 实例化聚合
func InstancePermissionAggregate(rootID string) *Factory {
	return &Factory{Root: domain.NewAggregateRoot(rootID)}
}

// InstanceOf 实例化聚合和领域服务
func (factory *Factory) InstanceOf() (ok bool) {
	if len(factory.Root.RootID) == 0 {
		return
	}
	factory.Application = &vo.Permission{}
	return true
}
