package clients

import (
	"user_context/rhombic/domain/package/vo"
)

type PackageClient interface {
	GetDetail(organizationID string) (vo vo.ValueObject)
}