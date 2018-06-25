package database

import (
	"github.com/cjburchell/survey/models"
	"gopkg.in/mgo.v2/bson"
)

const resultsCollection = "results"
const surveyCollection = "survey"

// Gets all the results for a given survey
func GetAllResults(surveyId string) (results []models.Result, err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	err = tempSession.DB(databaseName).C(resultsCollection).Find(bson.M{"surveyId": surveyId}).All(&results)
	return
}

// Gets the results for a given survey and question
func GetResults(surveyId string, questionId string) (results []models.Result, err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	err = tempSession.DB(databaseName).C(resultsCollection).Find(bson.M{"surveyId": surveyId, "questionId": questionId}).All(&results)
	return
}

// Increment a survey answer count
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

// Increment the survey submit count
func IncrementSubmitCount(surveyId string) (err error) {
	tempSession := session.Clone()
	defer tempSession.Close()
	query := bson.M{
		"surveyId": surveyId,
	}

	update := bson.M{
		"$inc": bson.M{
			"count": 1,
		},
	}

	_, err = tempSession.DB(databaseName).C(surveyCollection).Upsert(query, update)
	return
}

// Get the survey submit count
func GetSubmitCount(surveyId string) (count int) {
	tempSession := session.Clone()
	defer tempSession.Close()
	query := bson.M{
		"surveyId": surveyId,
	}

	var result struct {
		Count int `bson:"count"`
	}
	err := tempSession.DB(databaseName).C(surveyCollection).Find(query).One(&result)
	if err != nil {
		return 0
	}

	return result.Count
}
