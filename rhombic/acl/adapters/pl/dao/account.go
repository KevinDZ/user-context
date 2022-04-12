package dao

type Account struct{}

func (account *Account) TableName() string {
	return "account"
}

//dao 访问数据库

// RedisDAO redis数据库
type RedisDAO struct {
	Key string
}

// PostgreSQLDAO PostgreSQL数据库
type PostgreSQLDAO struct {
	ID string
}

// LoginRecordDAO 账户登录记录表
type LoginRecordDAO struct {
	ID int64
}
