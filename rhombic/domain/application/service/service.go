package service

import (
	port "user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/application/vo"
)

type Service struct {
	port.ApplicationClient
}

// GetList 应用列表
func (app *Service) GetList(id string) []vo.Application {
	return app.ApplicationClient.GetList(id)
}
