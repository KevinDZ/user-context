package clients

import (
	"context"
	"fmt"
	"log"

	"user-context/rhombic/acl/adapters/pl/dto"
	"user-context/rhombic/acl/ports/clients"
	"user-context/rhombic/domain/space/vo"
	"user-context/utils/grpc"
)

// SpaceAdapter 空间适配器，实现空间端口定义的方法
type SpaceAdapter struct {
	//Client *grpc.ClientConn
	SpaceServiceClient dto.SpaceServiceClient
}

// 检查是否实现了接口
var _ clients.SpaceClient = (*SpaceAdapter)(nil)

// NewSpaceAdapter 实现 grpc 客户端
func NewSpaceAdapter() clients.SpaceClient {
	return &SpaceAdapter{
		//Client: utils.NewClient(),
		SpaceServiceClient: dto.NewSpaceServiceClient(grpc.NewClient()),
	}
}

func (adapter *SpaceAdapter) GetList(organizationID string) (objects []vo.ValueObject) {
	// 1.请求参数
	request := &dto.SpaceRequest{OrganizationID: organizationID}
	// 2.proto 服务的方法
	respond, err := adapter.SpaceServiceClient.GetList(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
		return
	}
	fmt.Println("grpc respond: ", respond)

	// 领域模型值对象实例化
	objects = make([]vo.ValueObject, len(respond.GetSpace()))

	// grpc 转成 领域模型
	for key, value := range respond.GetSpace() {
		objects[key] = vo.ValueObject{
			Name:     value.GetName(),
			Owner:    value.GetOwner(),
			Manager:  value.GetManager(),
			Seat:     value.GetSeat(),
			Capacity: value.GetCapacity(),
		}
	}
	return
}
