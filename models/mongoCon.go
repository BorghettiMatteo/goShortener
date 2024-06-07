package models

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Client

type MongoCon struct{}

func (client *MongoCon) CreateDb() {

	dbString := os.Getenv("MONGODB_URL")
	if dbString == "" {
		panic("please specify a connection string goddamit!")
	}

	var err error
	Db, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI(
			dbString,
		),
	)
	if err != nil {
		panic("Not able to create mongodb connection: " + err.Error())
	}
	err = Db.Ping(context.TODO(), nil)
	if err != nil {
		println(err.Error())
	}
	// initialize collection
	Db.Database("db").Collection("pastamexs")

}

func (client *MongoCon) KillMongoDB() {
	Db.Disconnect(context.Background())
}
