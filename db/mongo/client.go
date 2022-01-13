package mongo

import (
	"context"
	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	db        *mongo.Database
	cli       *mongo.Client
	rawConfig client.Mongo
}

func NewClient(cli *mongo.Client, db *mongo.Database, rawConfig client.Mongo) *Client {
	return &Client{db: db, rawConfig: rawConfig, cli: cli}
}

//WatchHeartbeat 监测心跳和重连
func (c *Client) WatchHeartbeat() {
	//TODO 待实现
}

//Close 关闭连接
func (c *Client) Close() error {
	return c.cli.Disconnect(context.TODO())
}

//replaceDB 替换内部连接
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
