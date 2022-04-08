package pl
//dao 访问数据库

// RedisDAO redis数据库
type RedisDAO struct {
	Key string
}

// MysqlDAO MySQL数据库
type MysqlDAO struct {
	ID string
}

// LoginRecordDAO 账户登录记录表
type LoginRecordDAO struct {
	ID int64
	
}
