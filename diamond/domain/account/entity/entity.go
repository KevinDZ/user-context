package entity

// Account 账户实体
type Account struct {
	ID           string //global ID：账户ID，实体主键，Account内唯一
	NickName     string //账户名
	PassWord     string //账户密码
	Mobile       string //手机
	Email        string //邮箱
	ThirdPartyID string //第三方上下文 - ID: [微博ID、微信ID、今日头条ID...]
}
