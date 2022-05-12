package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ehwjh2010/viper/client/enums"
	"github.com/ehwjh2010/viper/client/settings"
)

// SetUp 初始化mongo
func SetUp(conf settings.Mongo) (*Client, error) {
	cli, db, err := setup(conf)
	if err != nil {
		return nil, err
	}
	c := NewClient(cli, db, conf)
	c.WatchHeartbeat()
	return c, nil
}

func setup(conf settings.Mongo) (*mongo.Client, *mongo.Database, error) {
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
