package service

import (
	adapter "user_context/rhombic/acl/adapters/clients"
	port "user_context/rhombic/acl/ports/clients"
	"user_context/rhombic/domain/application/vo"
)

type Service struct {
	port.ApplicationClient
}

func NewApplicationService() *Service {
	return &Service{
		adapter.NewApplicationAdapter(),
	}
}

// GetList 应用列表
func (app *Service)GetList(id string) []vo.ValueObject {
	return app.ApplicationClient.GetList(id)
}
