package publishers

import (
	"github.com/go-redis/redis"
	"sync"
	"user_context/rhombic/acl/ports/publishers"
)

// AccountEvent 账户发布者，实现账户端口定义的方法
type AccountEvent struct {
	// 发布者实体
	db *redis.IntCmd //TODO 具体调用方法查看文档
}

var (
	once sync.Once
	pub  publishers.AccountPublisher
)

func NewAccountEvent(channel string) publishers.AccountPublisher {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{

		})
		pub = &AccountEvent{
			db: client.Publish(channel, ""), // TODO 事件发布者的实体方法
		}
	})
	return pub
}

func (event AccountEvent)Registered(id, msg string)(ok bool){
	return
}

func (event AccountEvent)BindWechat(id, msg string)(err error){
	return
}