package db

import (
	"context"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	pb "emprpc/employee"
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
}

// GetDB returns the mongodb config instance
func GetDB() *MCfg {
	return mgo
}

// GetEmployee returns a document from the employee collection
func (m *MCfg) GetEmployee(f bson.M, emp *pb.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := m.MClient.Database(m.DB).Collection("employee").FindOne(ctx, f).Decode(&emp); err != nil {
		return err
	}
	spew.Dump(emp)
	return nil
}
