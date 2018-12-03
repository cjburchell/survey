package database

import (
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/tools-go/env"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

var databaseName string

// Connect to the database
func Connect() (err error) {
	databaseName = env.Get("DATABASE_NAME", "survey")
	databaseURL := env.Get("DATABASE_URL", "mongodb")
	log.Printf("Connecting to Database at %s", databaseURL)
	session, err = mgo.Dial(databaseURL)
	return
}
