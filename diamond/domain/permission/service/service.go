package service

import (
	"user-context/diamond/acl/ports/clients"
	"user-context/diamond/domain/permission/vo"
)

type Service struct {
	Permission clients.PermissionClient
}

// GetDetail 权限详情
func (service *Service) GetDetail(id string) vo.Permission {
	return service.Permission.GetDetail(id)
}
