package dao

import (
	"github.com/ndphu/fm/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (d *DAO) SaveOrUpdateFile(fo *model.File) (*model.File, error) {
	var err error
	if fo.Id.Hex() == "" {
		fo.Id = bson.NewObjectId()
		err = d.FileCollection().Insert(fo)
	} else {
		err = d.FileCollection().UpdateId(fo.Id, fo)
	}
	return fo, err
}

func (d *DAO) FileCollection() *mgo.Collection {
	return d.db.C("file")
}

func (d *DAO) FindRootFolder() (*model.File, error) {
	return d.FileFile(bson.M{"isRoot": true})
}

func (d *DAO) FileFile(query bson.M) (*model.File, error) {
	fo := model.File{}
	err := d.FileCollection().Find(query).One(&fo)

	return &fo, err
}
