package service

import (
	port "user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/permission/vo"
)

type Permission struct {
	port.PermissionClient
}

// GetDetail 权限详情
func (app *Permission) GetDetail(id string) vo.Permission {
	return app.PermissionClient.GetDetail(id)
}
