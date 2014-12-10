package ws

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func DbSave(collection string, data interface{}) error {
	session, err := mgo.Dial(dbUrl())

	if err != nil {
		return err
	}

	db := session.DB(dbName())

	defer session.Close()
	return db.C(collection).Insert(data)
}

func DbUpdate(collection string, query bson.M, data interface{}) error {
	session, err := mgo.Dial(dbUrl())

	if err != nil {
		return err
	}

	db := session.DB(dbName())

	defer session.Close()
	return db.C(collection).Update(query, data)
}

func DbFindOne(collection string, query bson.M, sort string, result interface{}) (err error) {
	session, err := mgo.Dial(dbUrl())

	if err != nil {
		return err
	}

	db := session.DB(dbName())

	defer session.Close()
	err = db.C(collection).Find(query).Sort(sort).One(result)

	return err
}

func dbUrl() string {
	url := os.Getenv("MONGO_URL")
	if url == "" {
		url = "mongodb://localhost"
	}
	return url
}

func dbName() string {
	db := os.Getenv("MONGO_DB")
	if db == "" {
		db = "reservations"
	}
	return db
}
