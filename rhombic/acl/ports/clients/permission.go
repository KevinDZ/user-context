package clients

import (
	"user_context/rhombic/domain/permission/vo"
)

type PermissionClient interface {
	GetDetail(organizationID string) (vo vo.ValueObject)
}