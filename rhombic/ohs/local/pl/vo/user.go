package vo

// view object 前端UI展示的对象

// RegisteredRequest 账户注册请求
type RegisteredRequest struct {
	RootID        string
	Mobile        string `json:"mobile"`             //手机号码
	MobileCaptcha string `json:"mobile_captcha"`     //验证码
	PassWord      string `json:"pass_word"`          // 密码
	NickName      string `json:"nickname,omitempty"` //TODO 昵称
	Trade         string `json:"trade"`              // 行业
	Position      string `json:"position"`           // 职位
	BandID        string `json:"band_id,omitempty"`  // 微信ID

	Account string `json:"account,omitempty"` //TODO 无账号
	Email   string `json:"email,omitempty"`   //TODO 无邮箱

	// header 信息
	IP         string `json:"ip"`
	ClientType string `json:"client_type,omitempty"`
}

// LoginRespond 账户登录响应
type LoginRespond struct {
	Code    int
	Message string
	Info    LoginInfo
}

type LoginInfo struct {
	ID        string `json:"id"`
	UID       string `json:"uid"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	AvatarUrl string `json:"avatar_url"`
	Company   string `json:"company"`
	Trade     string `json:"trade"`
	Position  string `json:"position"`
	AuthToken string `json:"auth_token"`
}

// LoginRequest 账户登录请求
type LoginRequest struct {
	RootID   string
	UserID   string `json:"user_id,omitempty"`
	Account  string `json:"account,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// LogoutRequest 账户登出请求
type LogoutRequest struct {
	ID string `json:"id,omitempty"`
}

// LogoutRespond 账户登出响应
type LogoutRespond struct {
}
