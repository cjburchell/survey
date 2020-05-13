package database

import (
	"github.com/cjburchell/settings-go"
	"github.com/cjburchell/uatu-go"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

var databaseName string

// Connect to the database
func Connect(log log.ILog, settings settings.ISettings) (err error) {
	databaseName = settings.Get("DATABASE_NAME", "survey")
	databaseURL := settings.Get("DATABASE_URL", "mongodb")
	log.Printf("Connecting to Database at %s", databaseURL)
	session, err = mgo.Dial(databaseURL)
	return
}
