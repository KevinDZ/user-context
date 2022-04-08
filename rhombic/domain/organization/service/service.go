package service

import (
	clientAdapter "user_context/rhombic/acl/adapters/clients"
	port "user_context/rhombic/acl/ports/clients"
	"user_context/rhombic/domain/organization/entity"
	"user_context/rhombic/domain/organization/vo"
)

type Service struct {
	port.OrganizationClient
}

func NewOrganizationService() *Service {
	return &Service{
		clientAdapter.NewOrganizationClient(),
	}
}

// GetList 获取组织的列表信息
func (organization *Service) GetList(organizationID string) []vo.ValueObject {
	return organization.OrganizationClient.GetList(organizationID)
}

// GetDetail 获取组织的列表信息
func (organization *Service) GetDetail(organizationID string) entity.Entity {
	return organization.OrganizationClient.GetDetail(organizationID)
}