package clients

import (
	"user-context/diamond/domain/application/vo"
)

type ApplicationClient interface {
	GetList(spaceID string) []vo.Application
}
