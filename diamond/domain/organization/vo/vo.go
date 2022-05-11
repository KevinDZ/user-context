package vo

import (
	space "user-context/diamond/domain/space/vo"
)

// Organization 值对象
type Organization struct {
	Person     []space.Space
	Enterprise []space.Space
}
