package service

import (
	port "user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/package/vo"
)

type Service struct {
	port.PackageClient
}

// GetList 应用列表
func (app *Service) GetList(id string) vo.Package {
	return app.PackageClient.GetDetail(id)
}
