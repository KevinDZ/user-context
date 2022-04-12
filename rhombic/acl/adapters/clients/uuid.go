package clients

import (
	"context"
	"fmt"
	"log"
	"user-context/rhombic/acl/adapters/pl/dto"
	"user-context/rhombic/acl/ports/clients"
	"user-context/utils/grpc"
)

// UUIDAdapter 雪花算法适配器，实现UUID端口定义的方法
type UUIDAdapter struct {
	//Client *grpc.ClientConn
	UUIDServiceClient dto.UUIDServiceClient
}

// 检查是否实现了接口
var _ clients.UUIDClient = (*UUIDAdapter)(nil)

func NewUUIDAdapter() clients.UUIDClient {
	return &UUIDAdapter{
		UUIDServiceClient: dto.NewUUIDServiceClient(grpc.NewClient()),
	}
}

func (adapter *UUIDAdapter) GetUUID() (uuid string) {
	request := &dto.UUIDRequest{}
	respond, err := adapter.UUIDServiceClient.GetUUID(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
		return
	}
	fmt.Println("clients.uuid grpc respond: ", respond)

	if len(respond.GetUUID()) == 0 {
		log.Fatalf("uuid get failed")
		return
	}
	return respond.GetUUID()
}
