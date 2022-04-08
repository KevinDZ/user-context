package clients

import (
	"user_context/rhombic/acl/ports/clients"
	"user_context/rhombic/domain/space/vo"
)

// SpaceAdapter 空间适配器，实现空间端口定义的方法
type SpaceAdapter struct {
	//grpc 方法调用
}


// 检查是否实现了接口
var _ clients.SpaceClient = (*SpaceAdapter)(nil)


func NewSpaceAdapter() clients.SpaceClient {
	return &SpaceAdapter{
		//grpc 方法调用
	}
}

func (adapter *SpaceAdapter) GetList(organizationID string) (vo []vo.ValueObject) {

	return
}