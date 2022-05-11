package publishers

// AccountPublisher 应用事件
type AccountPublisher interface {
	Registered(event map[string]string) error // 注册账户
	BindWechat(event map[string]string) error // 绑定微信
	Close() error
}
