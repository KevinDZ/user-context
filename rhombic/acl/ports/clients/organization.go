package clients

import (
	"user-context/rhombic/domain/organization/entity"
	"user-context/rhombic/domain/organization/vo"
)

type OrganizationClient interface {
	GetList(organizationID string) []vo.Organization
	GetDetail(organizationID string) entity.Organization
}
