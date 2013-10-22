package main

import (
	"labix.org/v2/mgo"
)

type Measurement struct {
	Weight float64
}

type Model struct {
	Database *mgo.Database
}

func Connect(uri string, database string) *Model {
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	model := new(Model)
	model.Database = session.DB(database)

	return model
}

func (model *Model) Close() {
	model.Database.Session.Close()
}

func (model *Model) InsertWeight(weight float64) (err error) {

	c := model.Database.C("measurements")
	err = c.Insert(&Measurement{weight})
	if err != nil {
		panic(err)
	}
	return
}
