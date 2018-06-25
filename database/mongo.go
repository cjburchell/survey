package database

import (
	"github.com/cjburchell/reefstatus-go/common"
	"github.com/cjburchell/reefstatus-go/common/log"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

var databaseName string

// Connect to the database
func Connect() (err error) {
	databaseName = common.GetEnv("DATABASE_NAME", "survey")
	databaseUrl := common.GetEnv("DATABASE_URL", "mongodb")
	log.Printf("Connecting to Database at %s", databaseUrl)
	session, err = mgo.Dial(databaseUrl)
	return
}
