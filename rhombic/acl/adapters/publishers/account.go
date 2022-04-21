package publishers

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
	"user-context/rhombic/acl/ports/publishers"
)

// AccountEvent 账户发布者，实现账户端口定义的方法
type AccountEvent struct {
	// TODO 发布者实体：redis or MQ
	*redis.Client //TODO 具体调用方法查看文档
}

var (
	once sync.Once
	pub  publishers.AccountPublisher
)

func NewAccountEvent() publishers.AccountPublisher {
	once.Do(func() {
		// TODO 事件发布者的实体方法 client.Publish(channel, "")
		pub = &AccountEvent{redis.NewClient(&redis.Options{})}
	})
	return pub
}

// Registered 注册事件，失败后重试机制的实现逻辑
func (event AccountEvent) Registered(channel string, msg map[string]string) (err error) {
	count := 5
RETRY:
	err = event.Publish(channel, msg).Err()
	if err != nil {
		for count < 5 {
			time.Sleep(time.Duration(2<<count) * time.Second)
			count++
			goto RETRY
		}
	}
	return
}

func (event AccountEvent) BindWechat(channel string, msg map[string]string) (err error) {
	count := 5
RETRY:
	err = event.Publish(channel, msg).Err()
	if err != nil {
		for count < 5 {
			time.Sleep(time.Duration(2<<count) * time.Second)
			count++
			goto RETRY
		}
	}
	return
}
