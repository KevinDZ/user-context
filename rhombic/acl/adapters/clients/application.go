package clients

import (
	"user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/application/vo"
)

// ApplicationAdapter 组织适配器，实现组织端口定义的方法
type ApplicationAdapter struct {
	//grpc 方法调用
}

// 检查是否实现了接口
var _ clients.ApplicationClient = (*ApplicationAdapter)(nil)

func NewApplicationAdapter() clients.ApplicationClient {

	return &ApplicationAdapter{
		//grpc 方法调用
	}
}

func (adapter *ApplicationAdapter) GetList(organizationID string) (list []vo.Application) {
	return
}
