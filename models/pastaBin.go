package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PastaBin struct {
	DbId        string    `bson:"_id,omitempty" json:",omitempty"`
	Body        []byte    `json:"body" bson:"body"`
	MessageType string    `json:"messagetype" bson:"messagetype,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty" bson:"timestamp.omitempty"`
}

func (p *PastaBin) InsertPastaMexToDb() (error, string) {
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//insert in mongodb
	res, err := Db.Database("db").Collection("pastamexs").InsertOne(ct, p)
	idToReturn := res.InsertedID.(primitive.ObjectID).Hex()
	return err, idToReturn
}
