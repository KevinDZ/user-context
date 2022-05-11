package service

import (
	port "user-context/diamond/acl/ports/clients"
	"user-context/diamond/domain/permission/vo"
)

type Permission struct {
	port.PermissionClient
}

// GetDetail 权限详情
func (app *Permission) GetDetail(id string) vo.Permission {
	return app.PermissionClient.GetDetail(id)
}
