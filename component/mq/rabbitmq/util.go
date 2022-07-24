package rabbitmq

import (
	"github.com/ehwjh2010/viper/helper/basic/str"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"

	"github.com/ehwjh2010/viper/log"
)

// Connect 连接
func Connect(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Error("connect rabbitmq failed", zap.Error(err), zap.String("Url", url))
		return nil, err
	}
	return conn, nil
}

// GetChannel 获取Channel
func GetChannel(conn *amqp.Connection) (*amqp.Channel, error) {

	if conn == nil {
		return nil, EmptyConnection
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

// CloseChannel 关闭Channel
func CloseChannel(ch *amqp.Channel) error {
	if ch == nil {
		return nil
	}

	return ch.Close()
}

// ExchangeDeclare 声明交换机
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

// QueueDeclare 声明队列
func QueueDeclare(ch *amqp.Channel, queue Queue) error {
	if _, err := ch.QueueDeclare(
		queue.Name,
		queue.Persistence,
		queue.AutoDeleted,
		queue.Exclusive,
		queue.NoWait,
		queue.Arguments,
	); err != nil {
		return err
	}

	return nil
}

// BindExchangeQueue 交换机绑定队列
func BindExchangeQueue(ch *amqp.Channel, queueName, exchangeName, routingKey string, broadcast bool) error {
	if !broadcast && str.IsEmpty(routingKey) {
		return EmptyRoutingKey
	}

	if err := ch.QueueBind(
		queueName,    // 绑定的队列名称
		routingKey,   // routingKey 用于消息路由分发的key
		exchangeName, // 绑定的exchange名
		false,        // 非阻塞
		nil,          // 额外属性
	); err != nil {
		return err
	}

	return nil
}
