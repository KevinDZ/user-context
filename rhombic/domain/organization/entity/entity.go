package entity

import (
	vo2 "user-context/rhombic/domain/package/vo"
	"user-context/rhombic/domain/permission/vo"
)

// Entity 实体
type Entity struct {
	ID            string            //空间ID，实体主键，CurrentSpace内唯一
	UserID        string            //成员ID
	Name          string            //空间名称
	Owner         string            //拥有者
	Collaborators map[string]string //协作者
	Seat          int               //席位
	Capacity      int               //容量
	Package       vo2.ValueObject   //套餐
	Permission    vo.ValueObject    //空间权限
}
