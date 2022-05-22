package mongo

import (
	"context"
	"github.com/ehwjh2010/viper/enums"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetUp 初始化mongo
func SetUp(conf Mongo) (*Client, error) {
	cli, db, err := setup(conf)
	if err != nil {
		return nil, err
	}
	c := NewClient(cli, db, conf)
	c.WatchHeartbeat()
	return c, nil
}

func setup(conf Mongo) (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), enums.FiveMinute)
	defer cancel()
	o := options.Client().ApplyURI(conf.Uri)
	o.SetMaxPoolSize(conf.MaxOpenConnCount)
	o.SetMinPoolSize(conf.MinOpenConnCount)
	o.SetMaxConnIdleTime(time.Duration(conf.FreeMaxLifetime) * time.Second)
	cli, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, nil, err
	}

	db := cli.Database(conf.Database)
	return cli, db, nil
}
