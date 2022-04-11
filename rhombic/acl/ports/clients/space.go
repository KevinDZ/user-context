package clients

import (
	"user-context/rhombic/domain/space/vo"
)

type SpaceClient interface {
	GetList(organizationID string) []vo.ValueObject
}
