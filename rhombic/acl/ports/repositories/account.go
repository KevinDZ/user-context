package repositories

import (
	"user-context/rhombic/domain/account/entity"
)

type AccountRepository interface {
	CheckIsExist(entity entity.Account) error
	Insert(entity entity.Account) error
	Query(id string) (entity *entity.Account)
	Update(entity entity.Account) error
	Delete(id string) error
	//save 持久化功能
}
