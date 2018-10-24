package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cjburchell/survey/models"
)

var surveys map[string]models.Survey

// GetSurvey gets a survey given the survey id
func GetSurvey(surveyID string) (*models.Survey, error) {
	// check to see if the survey is in the cash
	if survey, ok := surveys[surveyID]; ok {
		return &survey, nil
	}

	// load it from the json file
	raw, err := ioutil.ReadFile(fmt.Sprintf("survey%s.json", surveyID))
	if err != nil {
		return nil, err
	}

	var survey models.Survey
	err = json.Unmarshal(raw, &survey)
	return &survey, err
}
