package clients

import (
	"user-context/diamond/domain/package/vo"
)

type PackageClient interface {
	GetDetail(organizationID string) (vo vo.Package)
}
