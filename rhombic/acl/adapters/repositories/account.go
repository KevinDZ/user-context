package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	goRedis "github.com/go-redis/redis"
	"user-context/rhombic/acl/adapters/pl/dao"
	"user-context/rhombic/acl/ports/repositories"
	"user-context/rhombic/domain/account/entity"
	"user-context/rhombic/ohs/local/pl/vo"
	"user-context/utils/common/db"
	"user-context/utils/common/redis"
)

// TODO 数据库线程池

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

func (adapter *AccountAdapter) CheckIsExist(entity entity.Entity) (err error) {
	//查询数据库记录
	account := dao.Account{}
	err = adapter.db.First(&account, "id = ?", entity.ID).Error
	if account.ID == "" {
		err = errors.New("account no exist")
		return
	}
	return
}

func (adapter *AccountAdapter) Insert(entity entity.Entity) (err error) {
	// 领域模型转数据模型
	return adapter.db.Create(&dao.Account{
		// TODO
		ID: entity.ID,
	}).Error
}

func (adapter *AccountAdapter) Query(id string) (entities *entity.Entity) {
	var account dao.Account
	err := adapter.db.First(&account, "id = ?", id).Error
	if err != nil {
		return
	}
	entities = &entity.Entity{}
	return
}

func (adapter *AccountAdapter) Update(entity entity.Entity) (err error) {
	var account dao.Account
	err = adapter.db.First(&account, "id = ?", entity.ID).Error
	if err != nil {
		return
	}
	err = adapter.db.Model(account).Updates(map[string]interface{}{
		// TODO
		"nickname": entity.NickName,
	}).Error
	return
}

func (adapter *AccountAdapter) Delete(id string) (err error) {
	account := &dao.Account{ID: id}
	adapter.db.Delete(account)
	return
}

// DAO CQRS原则  读操作
type DAO struct {
	// 实例化数据库
	BD    *gorm.DB
	Redis *goRedis.Client
	// 数据模型
	User dao.UserDAO
}

// NewDAO 实例化数据库
func NewDAO(types string) *DAO {
	switch types {
	case "redis":
		return &DAO{Redis: redis.REDIS}
	case "db":
		return &DAO{BD: db.NewDBEngine()}
	default:
		return nil
	}
}

//LoginRecord 用户登录记录表 - 写操作
func (d *DAO) LoginRecord(request vo.LoginRequest) (err error) {
	err = d.BD.Create(&dao.LoginRecordDAO{
		// TODO
	}).Error
	return
}

// AccountDAO CQRS原则：账户数据访问对象  - 读操作
func (d *DAO) AccountDAO(userID string) (err error) {
	// 用户数据表 - 读操作
	err = d.BD.First(&d.User, "id = ?", userID).Error
	if err != nil {
		return
	}
	if d.User.UserID == "" { // 用户是否存在
		err = errors.New("user no exist")
	}
	return
}

// CreateToken TODO 创建token
func (d *DAO) CreateToken() (token string, err error) {
	// TODO set login ticket
	var expire int64
	switch d.User.ClientType {
	case "app":
		expire = time.Now().AddDate(0, 1, 0).Unix()
	case "web":
		expire = time.Now().AddDate(0, 0, 7).Unix()
	default:
		return
	}
	info := dao.UserInfo{
		ID:         d.User.UserID,
		UID:        d.User.UUID,
		UserName:   d.User.UUID,
		ClientType: d.User.ClientType,
		Nickname:   d.User.Nickname,
		Mobile:     d.User.Mobile,
		AvatarUrl:  d.User.AvatarUrl,
		SiteCode:   d.User.SiteCode,
	}
	request := dao.AuthRequest{
		ExpirationTime: expire,
		Info:           info,
	}
	requestJson, jsonErr := json.Marshal(request)
	if err != nil {
		err = jsonErr
		return
	}
	// auth-service 交互
	url := ""
	resp, respErr := http.NewRequest("POST", url, bytes.NewBuffer(requestJson))
	if err != nil {
		err = respErr
		return
	}
	defer resp.Body.Close()
	respBody, respErr := ioutil.ReadAll(resp.Body)
	if respErr != nil {
		err = respErr
		return
	}
	var respond dao.AuthRespond
	if err = json.Unmarshal(respBody, &respond); err != nil {
		return
	}
	token = respond.Data.Token
	return
}

func (d *DAO) MobileVerify(mobile, captcha string) (err error) {
	value := d.Redis.Get(fmt.Sprintf("%v%v", mobile, captcha))
	err = value.Err()
	if err != nil {
		return
	}
	if value.String() != captcha {
		err = errors.New("CaptchaCodeError")
		return
	}
	return
}

// EventMessage 记录事件消息
func (d *DAO) EventMessage() (err error) {
	err = d.BD.Create(&dao.LoginRecordDAO{
		// TODO
	}).Error
	return
}
