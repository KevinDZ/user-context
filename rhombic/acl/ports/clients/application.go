package clients

import (
	"user-context/rhombic/domain/application/vo"
)

type ApplicationClient interface {
	GetList(spaceID string) []vo.ValueObject
}
