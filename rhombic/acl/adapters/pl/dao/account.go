package dao

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
)

type Account struct{}

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

func (dao *RedisDAO) MobileVerify(mobile, captcha string) (err error) {
	// redis 实例化
	client := redis.NewClient(&redis.Options{})
	value := client.Get(fmt.Sprintf("%v%v", mobile, captcha))
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

// PostgreSQLDAO PostgreSQL数据库
type PostgreSQLDAO struct {
	ID string
}

// LoginRecordDAO 账户登录记录表
type LoginRecordDAO struct {
	ID int64
}
