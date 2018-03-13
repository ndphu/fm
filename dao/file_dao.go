package dao

import (
	"gopkg.in/mgo.v2"
)

func (d *DAO) FileCollection() *mgo.Collection {
	return d.db.C("file")
}
