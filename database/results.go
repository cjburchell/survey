package database

import (
	"github.com/cjburchell/survey/models"
	"gopkg.in/mgo.v2/bson"
)

const resultsCollection = "results"
const surveyCollection = "survey"

// GetAllResults gets all the results for a given survey
func GetAllResults(surveyID string) (results []models.Result, err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	err = tempSession.DB(databaseName).C(resultsCollection).Find(bson.M{"surveyId": surveyID}).All(&results)
	return
}

// GetResults gets the results for a given survey and question
func GetResults(surveyID string, questionId string) (results []models.Result, err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	err = tempSession.DB(databaseName).C(resultsCollection).Find(bson.M{"surveyId": surveyID, "questionId": questionId}).All(&results)
	return
}

// IncrementResult increments a survey answer count
func IncrementResult(surveyID string, questionId string, answer string) (err error) {
	tempSession := session.Clone()
	defer tempSession.Close()

	query := bson.M{
		"surveyId":   surveyID,
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

// IncrementSubmitCount increments the survey submit count
func IncrementSubmitCount(surveyID string) (err error) {
	tempSession := session.Clone()
	defer tempSession.Close()
	query := bson.M{
		"surveyId": surveyID,
	}

	update := bson.M{
		"$inc": bson.M{
			"count": 1,
		},
	}

	_, err = tempSession.DB(databaseName).C(surveyCollection).Upsert(query, update)
	return
}

// GetSubmitCount Gets the survey submit count
func GetSubmitCount(surveyID string) (count int) {
	tempSession := session.Clone()
	defer tempSession.Close()
	query := bson.M{
		"surveyId": surveyID,
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
