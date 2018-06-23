package database

import (
	"github.com/cjburchell/reefstatus-go/common"
	"github.com/cjburchell/reefstatus-go/common/log"
	"github.com/cjburchell/survey/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

var databaseName string

const resultsCollection = "results"

func Connect() (err error) {
	databaseName = common.GetEnv("DATABASE_NAME", "survey")
	databaseUrl := common.GetEnv("DATABASE_URL", "mongodb")
	log.Printf("Connecting to Database at %s", databaseUrl)
	session, err = mgo.Dial(databaseUrl)
	return
}

func GetAllResults() (results []models.SurveyResult, err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	err = tempSession.DB(databaseName).C(resultsCollection).Find(bson.M{}).All(&results)
	return
}

func GetResults(questionId string) (results []models.SurveyResult, err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	err = tempSession.DB(databaseName).C(resultsCollection).Find(bson.M{"questionId": questionId}).All(&results)
	return
}

func IncrementResult(questionId string, answer string) (err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	query := bson.M{
		"questionId": questionId,
		"answer":     answer,
	}

	update := bson.M{
		"$inc": bson.M{
			"count": 1,
		},
	}

	_, err = tempSession.DB(databaseName).C(resultsCollection).Upsert(query, update)
	return
}
