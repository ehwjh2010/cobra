package mongo

import "go.mongodb.org/mongo-driver/mongo"

var clients = make([]*mongo.Client, 0)
