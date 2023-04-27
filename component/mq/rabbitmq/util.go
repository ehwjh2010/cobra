package rabbitmq

import (
	"github.com/ehwjh2010/viper/helper/basic/collection"
	"github.com/ehwjh2010/viper/helper/basic/str"
	"github.com/ehwjh2010/viper/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

// Connect 连接.
func Connect(url string) (*amqp.Connection, error) {
	conn, err := amqp.DialConfig(url, amqp.Config{
		Heartbeat: time.Second * 5,
	})
	if err != nil {
		log.Errorf("connect rabbitmq failed, err: %s, url: %s", err, url)
		return nil, err
	}
	return conn, nil
}

// ExchangeDeclare 声明交换机.
func ExchangeDeclare(ch *amqp.Channel, exchange Exchange) error {
	// 当交换机name为空, 则使用默认交换机, 不需要声明
	if !str.IsEmpty(exchange.Name) {
		if err := ch.ExchangeDeclare(
			exchange.Name,
			exchange.ExType,
			exchange.Persistence,
			exchange.AutoDeleted,
			exchange.Internal,
			exchange.NoWait,
			exchange.Arguments); err != nil {
			return err
		}
	}
	return nil
}

// QueueDeclare 声明队列.
func QueueDeclare(ch *amqp.Channel, queue Queue) error {
	_, err := ch.QueueDeclare(
		queue.Name,
		queue.Persistence,
		queue.AutoDeleted,
		queue.Exclusive,
		queue.NoWait,
		queue.Arguments,
	)
	return err
}

// BindExchangeQueue 交换机绑定队列.
func BindExchangeQueue(
	ch *amqp.Channel,
	queueName,
	exchangeName string,
	routingKeys []string,
	exchangeType string) error {

	broadcast := exchangeType == Fanout
	noRoutingKeys := collection.IsEmptyStr(routingKeys)

	if !broadcast && noRoutingKeys {
		return EmptyRoutingKey
	}

	for _, routingKey := range routingKeys {
		if err := ch.QueueBind(
			queueName,    // 绑定的队列名称
			routingKey,   // routingKey 用于消息路由分发的key
			exchangeName, // 绑定的exchange名
			false,        // 非阻塞
			nil,          // 额外属性
		); err != nil {
			return err
		}
	}

	return nil
}
