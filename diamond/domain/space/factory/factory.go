package factory

import (
	"strings"
	"user-context/diamond/domain"
	"user-context/diamond/domain/space/service"
	"user-context/diamond/domain/space/vo"
)

// Factory 继承父类：聚合根隐性依赖
type Factory struct {
	Aggregate domain.AggregateRoot
	Space     vo.Space
	Service   *service.Service
}

// InstanceSpaceAggregate 实例化聚合
func InstanceSpaceAggregate(rootID string) *Factory {
	return &Factory{Aggregate: domain.NewAggregateRoot(rootID)}
}

// InstanceOf 实例化聚合和领域服务
func (factory *Factory) InstanceOf() (ok bool) {
	if strings.Replace(factory.Aggregate.GetAggregateRootID(), " ", "", -1) == "" {
		return
	}
	factory.Space = vo.Space{}
	return true
}
