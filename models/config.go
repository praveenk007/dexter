package models

import "gopkg.in/mgo.v2/bson"

type Config struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
}
