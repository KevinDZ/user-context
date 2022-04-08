package clients

import (
	"user_context/rhombic/domain/organization/entity"
	"user_context/rhombic/domain/organization/vo"
)

type OrganizationClient interface {
	GetList(organizationID string) []vo.ValueObject
	GetDetail(organizationID string) entity.Entity
}