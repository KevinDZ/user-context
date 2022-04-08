package service

import (
	adapter "user-context/rhombic/acl/adapters/clients"
	port "user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/permission/vo"
)

type Permission struct {
	port.PermissionClient
}

func NewPermissionService() *Permission {
	return &Permission{
		adapter.NewPermissionAdapter(),
	}
}

// GetDetail 权限详情
func (app *Permission) GetDetail(id string) vo.ValueObject {
	return app.PermissionClient.GetDetail(id)
}
