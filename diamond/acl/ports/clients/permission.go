package clients

import (
	"user-context/diamond/domain/permission/vo"
)

type PermissionClient interface {
	GetDetail(organizationID string) (vo vo.Permission)
}
