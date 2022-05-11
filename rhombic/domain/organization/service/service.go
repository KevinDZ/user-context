package service

import (
	port "user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/organization/entity"
	"user-context/rhombic/domain/organization/vo"
)

type Service struct {
	port.OrganizationClient
}

// GetList 获取组织的列表信息
func (organization *Service) GetList(organizationID string) []vo.Organization {
	return organization.OrganizationClient.GetList(organizationID)
}

// GetDetail 获取组织的列表信息
func (organization *Service) GetDetail(organizationID string) entity.Organization {
	return organization.OrganizationClient.GetDetail(organizationID)
}
