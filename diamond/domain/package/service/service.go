package service

import (
	"user-context/diamond/acl/ports/clients"
	"user-context/diamond/domain/package/vo"
)

type Service struct {
	Package clients.PackageClient
}

// GetList 应用列表
func (service *Service) GetList(id string) vo.Package {
	return service.Package.GetDetail(id)
}
