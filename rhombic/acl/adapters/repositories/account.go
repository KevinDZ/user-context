package repositories

import (
	"gorm.io/gorm"
	"sync"
	"user-context/rhombic/acl/adapters/pl/dao"
	"user-context/rhombic/acl/ports/repositories"
	"user-context/rhombic/domain/account/entity"
	"user-context/rhombic/ohs/local/pl/vo"
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

// DAO CQRS原则  读操作
type DAO struct {
	// 实例化数据库
	db *gorm.DB
}

func NewDAO() *DAO {
	return &DAO{db: db.NewDBEngine()}
}

//LoginRecord 用户登录记录表 - 写操作
func (d *DAO) LoginRecord(request vo.LoginRequest) (err error) {
	var dao dao.LoginRecordDAO
	_ = dao
	return
}

// AccountDAO CQRS原则：账户数据访问对象  - 读操作
func (d *DAO) AccountDAO(request vo.LoginRequest) (result vo.LoginRespond) {
	// 用户数据表 - 读操作
	dao := dao.PostgreSQLDAO{ID: request.UserID}
	_ = dao
	return
}

// IsExist redis判断用户是否存在
func (d *DAO) IsExist(request vo.LoginRequest) (ok bool) {
	// 1.通过 key 判断是否存在
	// 2.通过 value 判断是否正确
	return
}

// Delete 用户退出登录
func (d *DAO) Delete(request vo.LogoutRequest) (err error) {
	// 1.通过 key 删除
	return
}
