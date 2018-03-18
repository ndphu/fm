package dao

import (
	"gopkg.in/mgo.v2"
)

func (d *DAO) UserCollection() *mgo.Collection {
	return d.db.C("user")
}
