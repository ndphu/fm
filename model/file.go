package model

import (
	"gopkg.in/mgo.v2/bson"
)

type File struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name"`
	ServerPath string        `json:"serverPath" bson:"serverPath"`
	ParentId   bson.ObjectId `json:"parentId" bson:"parentId,omitempty"`
	Type       FileType      `json:"type"`
}

type FileType string

const (
	TYPE_FOLDER FileType = "FOLDER"
	TYPP_FILE   FileType = "FILE"
)
