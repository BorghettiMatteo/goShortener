package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PastaBin struct {
	//DbId        string    `bson:"_id,omitempty" json:",omitempty"`
	Body        string    `json:"body" bson:"body"`
	MessageType string    `json:"messagetype" bson:"messagetype,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty" bson:"timestamp.omitempty"`
}

func (p *PastaBin) InsertPastaMexToDb() (string, error) {
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//insert in mongodb
	res, err := Db.Database("db").Collection("pastamexs").InsertOne(ct, p)
	idToReturn := res.InsertedID.(primitive.ObjectID).Hex()
	return idToReturn, err
}
