package database

import (
	"github.com/cjburchell/survey/models"
	"gopkg.in/mgo.v2/bson"
)

const resultsCollection = "results"

func GetAllResults(surveyId string) (results []models.Result, err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	err = tempSession.DB(databaseName).C(resultsCollection).Find(bson.M{"surveyId": surveyId}).All(&results)
	return
}

func GetResults(surveyId string, questionId string) (results []models.Result, err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	err = tempSession.DB(databaseName).C(resultsCollection).Find(bson.M{"surveyId": surveyId, "questionId": questionId}).All(&results)
	return
}

func IncrementResult(surveyId string, questionId string, answer string) (err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	query := bson.M{
		"surveyId":   surveyId,
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
