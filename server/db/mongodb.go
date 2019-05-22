package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MCfg contains the configs for the mongodb
type MCfg struct {
	URI     string
	MClient *mongo.Client
	DB      string
}

var mgo *MCfg

// Init initializes the redis by populating the client
func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	mgo = &MCfg{
		URI: "mongodb://localhost:27017",
		DB:  "emprpc",
	}
	mc, err := mongo.Connect(ctx, options.Client().ApplyURI(mgo.URI))
	if err != nil {
		log.Fatal("Failed to connect to mongodb: ", err.Error())
	}
	mgo.MClient = mc
	mgo.MClient.Database(mgo.DB).Collection("employee")
}

// GetDB returns the mongodb database instance
func GetDB() *mongo.Database {
	return mgo.MClient.Database(mgo.DB)
}

// GetDBConfig returns the config instance for the mongodb
func GetDBConfig() *MCfg {
	return mgo
}
