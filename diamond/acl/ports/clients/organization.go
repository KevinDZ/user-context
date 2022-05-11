package clients

import (
	"user-context/diamond/domain/organization/entity"
	"user-context/diamond/domain/organization/vo"
)

type OrganizationClient interface {
	GetList(organizationID string) []vo.Organization
	GetDetail(organizationID string) entity.Organization
}
