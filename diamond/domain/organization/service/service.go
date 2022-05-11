package service

import (
	port "user-context/diamond/acl/ports/clients"
	"user-context/diamond/domain/organization/entity"
	"user-context/diamond/domain/organization/vo"
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
