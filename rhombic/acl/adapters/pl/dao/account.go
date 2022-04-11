package dao

type Account struct {
}

func (account *Account) TableName() string {
	return "account"
}
