package pl

type Account struct {

}

func (account *Account) TableName() string {
	return "account"
}