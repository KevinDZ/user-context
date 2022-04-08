package service

import (
	adapter "user-context/rhombic/acl/adapters/clients"
	port "user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/package/vo"
)

type Service struct {
	port.PackageClient
}

func NewPackageService() *Service {
	return &Service{
		adapter.NewPackageAdapter(),
	}
}

// GetList 应用列表
func (app *Service) GetList(id string) vo.ValueObject {
	return app.PackageClient.GetDetail(id)
}
