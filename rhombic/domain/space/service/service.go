package service

import (
	adapter "user_context/rhombic/acl/adapters/clients"
	port "user_context/rhombic/acl/ports/clients"
	"user_context/rhombic/domain/space/vo"
)

type Service struct {
	port.SpaceClient
}

func NewSpaceService() *Service {
	return &Service{
		adapter.NewSpaceAdapter(),
	}
}

// GetList 空间列表
func (app *Service)GetList(id string) []vo.ValueObject {
	return app.SpaceClient.GetList(id)
}