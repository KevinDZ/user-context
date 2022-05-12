package service

import (
	"user-context/diamond/acl/ports/clients"
	"user-context/diamond/domain/space/vo"
)

type Service struct {
	Space clients.SpaceClient
}

// GetList 空间列表
func (service *Service) GetList(id string) []vo.Space {
	return service.Space.GetList(id)
}
