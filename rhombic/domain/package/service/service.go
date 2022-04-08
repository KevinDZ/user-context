package service

import (
	adapter "user_context/rhombic/acl/adapters/clients"
	port "user_context/rhombic/acl/ports/clients"
	"user_context/rhombic/domain/package/vo"
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
func (app *Service)GetList(id string) vo.ValueObject {
	return app.PackageClient.GetDetail(id)
}