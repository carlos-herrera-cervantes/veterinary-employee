package db

import (
	"context"
	"log"
	"sync"
	"veterinary-employee/settings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *mongo.Database
}

func initDB() {
	uri := settings.InitializeMongoDB().Host
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalf("[initDB][error:%s]", err.Error())
	}

	data = &Data{
		DB: client.Database(settings.InitializeMongoDB().DefaultDB),
	}
}

func New() *Data {
	once.Do(initDB)
	return data
}
