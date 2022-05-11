package entity

import (
	packages "user-context/diamond/domain/package/vo"
	permission "user-context/diamond/domain/permission/vo"
)

// Organization 实体
type Organization struct {
	ID            string                //空间ID，实体主键，CurrentSpace内唯一
	UserID        string                //成员ID
	Name          string                //空间名称
	Owner         string                //拥有者
	Collaborators map[string]string     //协作者
	Seat          int                   //席位
	Capacity      int                   //容量
	Package       packages.Package      //套餐
	Permission    permission.Permission //空间权限
}
