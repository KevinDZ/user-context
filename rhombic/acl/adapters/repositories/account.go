package repositories

import (
	"gorm.io/gorm"
	"sync"
	"user-context/rhombic/acl/ports/repositories"
	"user-context/rhombic/domain/account/entity"
	"user-context/utils/common/db"
)

// AccountAdapter 账户适配器，实现账户端口定义的方法
type AccountAdapter struct {
	db *gorm.DB
}

var (
	once sync.Once
	repo repositories.AccountRepository
)

// 检查是否实现了接口
var _ repositories.AccountRepository = (*AccountAdapter)(nil)

func NewAccountAdapter() repositories.AccountRepository {
	once.Do(func() {
		repo = &AccountAdapter{
			// 创建数据库引擎
			db: db.NewDBEngine(),
		}
	})
	return repo
}

func (adapter *AccountAdapter) CheckIsExist(entity entity.Entity) (ok bool) {
	//查询数据库记录
	return
}

func (adapter *AccountAdapter) Insert(entity entity.Entity) (err error) {
	return
}

func (adapter *AccountAdapter) Query(id string) (entity entity.Entity) {
	return
}

func (adapter *AccountAdapter) Update(entity entity.Entity) (err error) {
	return
}

func (adapter *AccountAdapter) Delete(id string) (err error) {
	return
}

func (adapter *AccountAdapter) Save() *AccountAdapter {
	return &AccountAdapter{adapter.db.Save("")}
}
