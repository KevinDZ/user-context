package clients

import (
	"user_context/rhombic/acl/ports/clients"
	"user_context/rhombic/domain/organization/entity"
	"user_context/rhombic/domain/organization/vo"
)

// OrganizationAdapter 组织适配器，实现组织端口定义的方法
type OrganizationAdapter struct {
	//grpc 方法调用
}

// 检查是否实现了接口
var _ clients.OrganizationClient = (*OrganizationAdapter)(nil)

func NewOrganizationClient() clients.OrganizationClient {
	return &OrganizationAdapter{
			//grpc 方法调用
		}
}

func (adapter *OrganizationAdapter) GetList(organizationID string)(list []vo.ValueObject){
	return
}

func (adapter *OrganizationAdapter) GetDetail(organizationID string)(entity entity.Entity){
	return
}