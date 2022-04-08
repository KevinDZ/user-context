package vo

import (
	"user-context/rhombic/domain/space/vo"
)

// ValueObject 值对象
type ValueObject struct {
	Person     []vo.ValueObject
	Enterprise []vo.ValueObject
}
