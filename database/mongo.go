package database

import (
	"github.com/cjburchell/tools-go"
	"github.com/cjburchell/yasls-client-go"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

var databaseName string

// Connect to the database
func Connect() (err error) {
	databaseName = tools.GetEnv("DATABASE_NAME", "survey")
	databaseURL := tools.GetEnv("DATABASE_URL", "mongodb")
	log.Printf("Connecting to Database at %s", databaseURL)
	session, err = mgo.Dial(databaseURL)
	return
}
