package vo

import (
	space "user-context/rhombic/domain/space/vo"
)

// Organization 值对象
type Organization struct {
	Person     []space.Space
	Enterprise []space.Space
}
