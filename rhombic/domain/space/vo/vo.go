package vo

// ValueObject 值对象
type ValueObject struct{
	Name string//空间名称
	Owner string//拥有者
	Manager bool//管理者
	Seat int//席位
	Capacity int//容量
}
