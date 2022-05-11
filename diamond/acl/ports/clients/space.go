package clients

import (
	"user-context/diamond/domain/space/vo"
)

type SpaceClient interface {
	GetList(organizationID string) []vo.Space
}
