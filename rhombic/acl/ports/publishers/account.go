package publishers

// AccountPublisher 应用事件
type AccountPublisher interface {
	Registered(channel string, event map[string]string) bool  // 注册账户
	BindWechat(channel string, event map[string]string) error // 绑定微信
}
