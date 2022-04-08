package clients

import (
	"user_context/rhombic/domain/application/vo"
)

type ApplicationClient interface {
	GetList(spaceID string) []vo.ValueObject
}