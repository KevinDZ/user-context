package publishers

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"user-context/diamond/acl/ports/publishers"
)

// AccountEvent 账户发布者，实现账户端口定义的方法
type AccountEvent struct {
	// rabbitMQ 生产者
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

//var (
//	once sync.Once
//	pub  publishers.AccountPublisher
//)

func NewAccountEvent(name string) publishers.AccountPublisher {
	// rabbitMQ 连接
	mq := new(AccountEvent)
	var err error
	mq.conn, err = amqp.Dial(viper.GetString("rabbitMQ.url"))
	if err != nil {
		return nil
	}
	mq.ch, err = mq.conn.Channel()
	if err != nil {
		return nil
	}
	mq.q, err = mq.ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	return mq
}

func (event AccountEvent) ChannelClose() (err error) {
	return event.ch.Close()
}

func (event AccountEvent) Close() (err error) {
	return event.conn.Close()
}

func (event AccountEvent) send(exchange string, msg map[string]string) (err error) {
	var data []byte
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	return event.ch.Publish(
		exchange,
		event.q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})
}

// Registered 注册事件，失败后重试机制的实现逻辑
func (event AccountEvent) Registered(msg map[string]string) (err error) {
	return event.send(viper.GetString("exchange.user"), msg)
}

func (event AccountEvent) BindWechat(msg map[string]string) (err error) {
	return event.send(viper.GetString("exchange.wechat"), msg)
}
