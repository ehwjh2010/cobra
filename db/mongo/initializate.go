package mongo

import (
	"context"
	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/global"
	"github.com/ehwjh2010/viper/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//SetUp 初始化mongo
func SetUp(conf client.Mongo) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), global.FiveMinute)
	defer cancel()
	o := options.Client().ApplyURI(conf.Uri)
	o.SetMaxPoolSize(conf.MaxConnectCount)
	o.SetMinPoolSize(conf.MinFreeConnCount)
	o.SetMaxConnIdleTime(time.Duration(conf.FreeMaxLifetime) * time.Minute)
	cli, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	clients = append(clients, cli)

	db := cli.Database(conf.Database)

	return db, nil
}

func Close() error {
	if len(clients) <= 0 {
		return nil
	}

	var multiErr types.MultiErr

	for _, cli := range clients {
		multiErr.AddErr(cli.Disconnect(context.TODO()))
	}

	return multiErr.AsStdErr()
}
