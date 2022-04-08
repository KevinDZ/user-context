package clients

import (
	"user-context/rhombic/domain/permission/vo"
)

type PermissionClient interface {
	GetDetail(organizationID string) (vo vo.ValueObject)
}
