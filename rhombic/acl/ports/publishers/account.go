package publishers

// AccountPublisher 应用事件
type AccountPublisher interface {
	Registered(id, event string) bool// 注册账户
	BindWechat(id, event string) error// 绑定微信
}