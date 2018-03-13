package dao

import (
	"fmt"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var (
	SERVER = "192.168.99.100"
)

type DAO struct {
	session *mgo.Session
	db      *mgo.Database
}

func NewDB() *DAO {
	return &DAO{}
}

func (d *DAO) Init() error {
	fmt.Printf("Connecting to %s...\n", SERVER)
	s, err := mgo.Dial(SERVER)
	if err != nil {
		return err
	}
	d.session = s
	d.db = s.DB("fm")
	return nil
}

func (d *DAO) Shutdown() error {
	fmt.Println("Shutting down DB...")
	if d.session != nil {
		d.session.Close()
	}
	return nil
}
