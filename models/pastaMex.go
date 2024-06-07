package models

import (
	"time"
)

type PastaMex struct {
	DbId        string    `bson:"_id" json:",omitempty"`
	Body        string    `json:"body"`
	MessageType string    `json:"messagetype"`
	Timestamp   time.Time `json:"timestamp"`
}

func insertPastaMexToDb(mex string) {

}
