package connection

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
)

func ConnectionString() string {
	authParams := env.ConstantValues.MongoAuthParams
	dbName := env.ConstantValues.MongoDatabase
	host := env.MongoHost()
	password := env.MongoPassword()
	user := env.MongoUsername()

	connectionString := "mongodb://" + user + ":" + password + "@" + host + "/" + dbName + "?" + authParams

	return connectionString
}

func MongoCollection() *mongo.Collection {
	mongoOptions := options.Client().ApplyURI(ConnectionString())
	client, err := mongo.Connect(context.Background(), mongoOptions)
	if err != nil {
		logrus.Error("mongo client failed", err)
	}
	if err != nil {
		logrus.Error("mongo client connection failed", err)
	}
	collection := client.Database(env.ConstantValues.MongoDatabase).Collection(env.ConstantValues.MongoCollection)

	return collection
}
