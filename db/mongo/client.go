package mongo

import (
	"context"
	"errors"
	"github.com/ehwjh2010/viper/client/enums"
	"github.com/ehwjh2010/viper/client/settings"
	"github.com/ehwjh2010/viper/component/routine"
	"github.com/ehwjh2010/viper/log"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Client struct {
	db        *mongo.Database
	cli       *mongo.Client
	rawConfig settings.Mongo // 数据库配置配置
	pCount    int            // 心跳连续失败次数
	rCount    int            // 重连连续失败次数
}

func NewClient(cli *mongo.Client, db *mongo.Database, rawConfig settings.Mongo) *Client {
	return &Client{db: db, rawConfig: rawConfig, cli: cli}
}

// Heartbeat ping连接
func (c *Client) Heartbeat() error {
	return c.cli.Ping(context.TODO(), nil)
}

// WatchHeartbeat 监测心跳和重连
func (c *Client) WatchHeartbeat() {
	//TODO 监测逻辑接口化

	fn := func() {
		waitFlag := true
		for {
			if waitFlag {
				<-time.After(3 * time.Second)
			}

			//重连失败次数大于0, 直接重连
			if c.rCount > 0 {
				if c.rCount >= 3 {
					<-time.After(enums.OneSecDur)
				}
				if ok, _ := c.replaceDB(); ok {
					c.rCount = 0
					c.pCount = 0
					waitFlag = true
				} else {
					c.rCount++
					c.pCount++
					waitFlag = false
				}
				continue
			}

			if c.Heartbeat() != nil {
				c.pCount++
				//心跳连续3次失败, 触发重连
				if c.pCount >= 3 {
					if ok, _ := c.replaceDB(); ok {
						c.rCount = 0
						c.pCount = 0
						waitFlag = true
					} else {
						c.rCount++
						waitFlag = false
					}
				}
			} else {
				c.rCount = 0
				c.pCount = 0
				waitFlag = true
			}
		}
	}

	//优先使用协程池监听, 如果没有使用原生协程监听
	err := routine.AddTask(fn)
	if err != nil {
		if errors.Is(err, routine.NoEnableRoutinePool) {
			go fn()
		} else {
			log.Warn("watch heartbeat failed")
		}

	}
}

// Close 关闭连接
func (c *Client) Close() error {
	return c.cli.Disconnect(context.TODO())
}

// replaceDB 替换内部连接
func (c *Client) replaceDB() (bool, error) {
	cli, db, err := setup(c.rawConfig)
	if err != nil {
		log.Err("reconnect mongo failed", err)
		return false, err
	}

	//关闭之前的连接
	c.Close()

	c.db = db
	c.cli = cli
	return true, nil
}
