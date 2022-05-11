package clients

import (
	"user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/permission/vo"
)

// PermissionAdapter 组织适配器，实现组织端口定义的方法
type PermissionAdapter struct {
	//grpc 方法调用
}

// 检查是否实现了接口
var _ clients.PermissionClient = (*PermissionAdapter)(nil)

func NewPermissionAdapter() clients.PermissionClient {
	return &PermissionAdapter{
		//grpc 方法调用
	}
}

func (adapter *PermissionAdapter) GetDetail(organizationID string) (vo vo.Permission) {

	return
}
