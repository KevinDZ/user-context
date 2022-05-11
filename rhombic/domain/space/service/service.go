package service

import (
	port "user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/space/vo"
)

type Service struct {
	port.SpaceClient
}

// GetList 空间列表
func (app *Service) GetList(id string) []vo.Space {
	return app.SpaceClient.GetList(id)
}
