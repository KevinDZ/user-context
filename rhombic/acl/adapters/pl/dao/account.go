package dao

import (
	"github.com/go-redis/redis"
)

type Account struct {
	ID string
}

func (account *Account) TableName() string {
	return "account"
}

//dao 访问数据库

// RedisDAO redis数据库
type RedisDAO struct {
	Dao *redis.Client
}

func NewRedisDAO() *RedisDAO {
	return &RedisDAO{
		Dao: redis.NewClient(&redis.Options{}),
	}
}

// UserDAO 用户数据表
type UserDAO struct {
	UserID     string
	UUID       string
	ClientType string
	Nickname   string
	Mobile     string
	Email      string
	AvatarUrl  string
	SiteCode   string
}

// LoginRecordDAO 账户登录记录表
type LoginRecordDAO struct {
	ID string

	Status string // 状态： try、conform、cancel
}
