package service

import (
	"user-context/diamond/acl/ports/clients"
	"user-context/diamond/domain/organization/entity"
	"user-context/diamond/domain/organization/vo"
)

type Service struct {
	Organization clients.OrganizationClient
}

// GetList 获取组织的列表信息
func (service *Service) GetList(organizationID string) []vo.Organization {
	return service.Organization.GetList(organizationID)
}

// GetDetail 获取组织的列表信息
func (service *Service) GetDetail(organizationID string) entity.Organization {
	return service.Organization.GetDetail(organizationID)
}
