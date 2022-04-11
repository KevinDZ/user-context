package service

import (
	adapter "user-context/rhombic/acl/adapters/clients"
	port "user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/space/vo"
)

type Service struct {
	port.SpaceClient
}

// NewSpaceService 实例化空间的领域服务
func NewSpaceService() *Service {
	// 实现端口方法的适配器
	return &Service{adapter.NewSpaceAdapter()}
}

// GetList 空间列表
func (app *Service) GetList(id string) []vo.ValueObject {
	return app.SpaceClient.GetList(id)
}
