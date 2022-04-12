package entity

// Entity 账户实体
type Entity struct {
	ID           string `json:"id,omitempty"`             //global ID：账户ID，实体主键，Account内唯一
	Name         string `json:"name,omitempty"`           //账户名
	PassWord     string `json:"pass_word,omitempty"`      //账户密码
	Phone        string `json:"phone,omitempty"`          //手机
	Email        string `json:"email,omitempty"`          //邮箱
	ThirdPartyID string `json:"third_party_id,omitempty"` //第三方上下文 - ID: [微博ID、微信ID、今日头条ID...]
}
