package rabbitmq

import (
	"errors"

	"go.uber.org/zap"

	"github.com/ehwjh2010/viper/log"
)

var (
	EmptyConnection           = errors.New("empty connection get channel")
	EmptyExchange             = errors.New("empty exchange")
	EmptyQueue                = errors.New("empty queue")
	EmptyRoutingKeys          = errors.New("empty routing key")
	ExecuteStartBeforeSendMsg = errors.New("execute start function before send msg")
	ExecuteStartBeforeRecvMsg = errors.New("execute start function before receive msg")
)

type ConsumeCallback func(msgId string, body []byte, err error) error

type Consumer interface {
	Consume([]byte) error
	ConsumeSuccessCallback(msgId string, body []byte)
	ConsumeFailConsumeCallback(msgId string, body []byte, err error)
}

const (
	DefaultPullBatchSize = 10
)

func DefaultSuccessCallback(body []byte) error {
	log.Info("send rabbitmq msg success", zap.ByteString("body", body))
	return nil
}

func DefaultFailCallback(body []byte) error {
	log.Info("send rabbitmq msg fail", zap.ByteString("body", body))
	return nil
}

const (
	Direct  = "direct"
	Fanout  = "fanout"
	Topic   = "topic"
	Headers = "headers"
)
