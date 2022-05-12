package entity

// Account 账户实体
type Account struct {
	id           string //global ID：账户ID，实体主键，Account内唯一
	nickName     string //账户名
	passWord     string //账户密码
	mobile       string //手机
	email        string //邮箱
	thirdPartyID string //第三方上下文 - ID: [微博ID、微信ID、今日头条ID...]
}

func (account Account) GetID() string {
	return account.id
}

func (account Account) GetNickName() string {
	return account.nickName
}

func (account Account) GetPassWord() string {
	return account.passWord
}

func (account Account) GetMobile() string {
	return account.mobile
}

func (account Account) GetEmail() string {
	return account.email
}

func (account Account) GetThirdPartyID() string {
	return account.thirdPartyID
}

func (account Account) SetID(field string) {
	account.id = field
	return
}

func (account Account) SetNickName(field string) {
	account.id = field
	return
}

func (account Account) SetPassWord(field string) {
	account.passWord = field
	return
}

func (account Account) SetMobile(field string) {
	account.mobile = field
	return
}

func (account Account) SetEmail(field string) {
	account.email = field
	return
}

func (account Account) SetThirdPartyID(field string) {
	account.thirdPartyID = field
	return
}
