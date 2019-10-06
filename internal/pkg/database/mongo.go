package database

import (
	"context"
	"log"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	client           *mongo.Client
	users            *mongo.Collection
	films            *mongo.Collection
	usersAuthCookies *mongo.Collection
}

func InitMongo() *MongoConnection {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.MongoDbUri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	conn := MongoConnection{
		client:           client,
		users:            client.Database(configs.MongoDbName).Collection(configs.UsersCollectionName),
		films:            client.Database(configs.MongoDbName).Collection(configs.FilmsCollectionName),
		usersAuthCookies: client.Database(configs.MongoDbName).Collection(configs.CoockiesCollectionName),
	}
	return &conn
}
