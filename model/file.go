package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type File struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name"`
	ServerPath string        `json:"serverPath" bson:"serverPath"`
	IsRoot     bool          `json:"isRoot" bson:"isRoot"`
	Parent     mgo.DBRef
}
