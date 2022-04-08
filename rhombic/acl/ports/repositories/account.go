package repositories

import (
	"user-context/rhombic/domain/account/entity"
)

type AccountRepository interface {
	CheckIsExist(entity entity.Entity) bool
	Insert(entity entity.Entity) error
	Query(id string) (entity entity.Entity)
	Update(entity entity.Entity)
	Delete(id string) error
}
