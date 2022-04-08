package clients

import (
	"user_context/rhombic/acl/ports/clients"
	"user_context/rhombic/domain/package/vo"
)

// PackageAdapter 组织适配器，实现组织端口定义的方法
type PackageAdapter struct {
	//grpc 方法调用
}


// 检查是否实现了接口
var _ clients.PackageClient = (*PackageAdapter)(nil)


func NewPackageAdapter() clients.PackageClient {
	return &PackageAdapter{
		//grpc 方法调用
	}
}

func (adapter *PackageAdapter) GetDetail(organizationID string) (vo vo.ValueObject) {

	return
}