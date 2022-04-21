package dao

type UserInfo struct {
	ID         string `json:"id"`
	UID        string `json:"uid"`
	UserName   string `json:"user_name"`
	ClientType string `json:"client_type"`
	Nickname   string `json:"nickname"`
	Mobile     string `json:"mobile"`
	AvatarUrl  string `json:"avatar_url"`
	SiteCode   string `json:"site_code"`
}

type AuthRequest struct {
	ExpirationTime int64
	Info           UserInfo
}

type AuthRespond struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Token string `json:"token"`
	JwtID string `json:"jwt_id"`
}
