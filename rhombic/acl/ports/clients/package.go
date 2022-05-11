package clients

import (
	"user-context/rhombic/domain/package/vo"
)

type PackageClient interface {
	GetDetail(organizationID string) (vo vo.Package)
}
