package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	FirstName  string        `json:"firstName" bson:"firstName"`
	LastName   string        `json:"lastName" bson:"lastName"`
	ExternalId string        `json:"externalId" bson:"externalId"`
}
