package mongo

import (
	"context"
	"github.com/ehwjh2010/viper/client/enum"
	"github.com/ehwjh2010/viper/client/setting"
	"github.com/ehwjh2010/viper/client/verror"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//SetUp 初始化mongo
func SetUp(conf *setting.Mongo) (*Client, error) {
	if conf == nil {
		return nil, verror.InvalidConfig
	}

	cli, db, err := setup(conf)
	if err != nil {
		return nil, err
	}
	c := NewClient(cli, db, conf)
	c.WatchHeartbeat()
	return c, nil
}

func setup(conf *setting.Mongo) (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), enum.FiveMinute)
	defer cancel()
	o := options.Client().ApplyURI(conf.Uri)
	o.SetMaxPoolSize(conf.MaxConnectCount)
	o.SetMinPoolSize(conf.MinConnectCount)
	o.SetMaxConnIdleTime(time.Duration(conf.FreeMaxLifetime) * time.Minute)
	cli, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, nil, err
	}

	db := cli.Database(conf.Database)
	return cli, db, nil
}
