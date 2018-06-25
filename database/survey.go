package database

import (
	"encoding/json"
	"fmt"
	"github.com/cjburchell/survey/models"
	"io/ioutil"
)

var surveys map[string]models.Survey

// Get a survey given the survey id
func GetSurvey(surveyId string) (*models.Survey, error) {
	// check to see if the survey is in the cash
	if survey, ok := surveys[surveyId]; ok {
		return &survey, nil
	}

	// load it from the json file
	raw, err := ioutil.ReadFile(fmt.Sprintf("survey%s.json", surveyId))
	if err != nil {
		return nil, err
	}

	var survey models.Survey
	err = json.Unmarshal(raw, &survey)
	return &survey, err
}
