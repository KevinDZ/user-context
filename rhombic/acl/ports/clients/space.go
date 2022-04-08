package clients

import (
	"user_context/rhombic/domain/space/vo"
)

type SpaceClient interface {
	GetList(organizationID string) (vo []vo.ValueObject)
}