package vo

// view object 前端UI展示的对象

// RegisteredRequest 账户注册请求
type RegisteredRequest struct {
	Account  string `json:"account,omitempty"`
	PassWord string `json:"pass_word,omitempty"`
	Name     string `json:"name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Type     string `json:"type,omitempty"`
}

// RegisteredRespond 账户注册响应
type RegisteredRespond struct {
}

// LoginRequest 账户登录请求
type LoginRequest struct {
	ID string `json:"id,omitempty"`
}

// LoginRespond 账户登录响应
type LoginRespond struct {
}

// LogoutRequest 账户登出请求
type LogoutRequest struct {
	ID string `json:"id,omitempty"`
}

// LogoutRespond 账户登出响应
type LogoutRespond struct {
}
