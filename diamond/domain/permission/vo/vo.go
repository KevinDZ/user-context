package vo

// Permission 值对象
type Permission struct {
	Name         string //权限名称
	Level        string //权限等级
	BoundContext string //权限限界上下文
}
