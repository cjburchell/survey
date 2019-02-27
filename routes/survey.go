package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/survey/database"
	"github.com/cjburchell/survey/models"
	"github.com/gorilla/mux"
)

// Setup the routes
func Setup(router *mux.Router) {
	surveyRoute := router.PathPrefix("/survey").Subrouter()
	surveyRoute.HandleFunc("/{surveyId}", handleGetSurvey).Methods("GET")
	surveyRoute.HandleFunc("/{surveyId}/results", handleGetResults).Methods("GET")
	surveyRoute.HandleFunc("/{surveyId}/results/{questionId}", handleGetResultsForQuestion).Methods("GET")
	surveyRoute.HandleFunc("/{surveyId}/answers", handleSetAnswers).Methods("POST")
	surveyRoute.HandleFunc("/{surveyId}/count", handleGetSurveyCount).Methods("GET")

	router.HandleFunc("/@status", func(writer http.ResponseWriter, _ *http.Request) {
		reply, _ := json.Marshal("Ok")
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(reply)
	}).Methods("GET")
}

func handleGetSurvey(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleGetQuestions %s", request.URL.String())

	vars := mux.Vars(request)
	surveyID := vars["surveyId"]

	survey, err := database.GetSurvey(surveyID)
	if err != nil {
		log.Error(err, "Unable to get survey")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	reply, err := json.Marshal(survey)
	if err != nil {
		log.Error(err, "Marshal Failed")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(reply)
}

func handleGetSurveyCount(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleGetSurveyCount %s", request.URL.String())

	vars := mux.Vars(request)
	surveyID := vars["surveyId"]

	count := database.GetSubmitCount(surveyID)

	result := struct {
		Count int `json:"count"`
	}{
		count,
	}

	reply, err := json.Marshal(result)
	if err != nil {
		log.Error(err, "Marshal Failed")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(reply)
}

func handleGetResults(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleGetResults %s", request.URL.String())

	vars := mux.Vars(request)
	surveyID := vars["surveyId"]

	results, err := database.GetAllResults(surveyID)
	if err != nil {
		log.Error(err, "GetAllResults Failed")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if results == nil {
		writer.WriteHeader(http.StatusOK)
		return
	}

	reply, err := json.Marshal(results)
	if err != nil {
		log.Error(err, "Marshal Failed")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(reply)
}

func handleGetResultsForQuestion(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleGetResultsForId %s", request.URL.String())

	vars := mux.Vars(request)
	questionID := vars["questionId"]
	surveyID := vars["surveyId"]

	results, err := database.GetResults(surveyID, questionID)
	if err != nil {
		log.Error(err, "GetAllResults Failed")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if results == nil {
		writer.WriteHeader(http.StatusOK)
		return
	}

	reply, err := json.Marshal(results)
	if err != nil {
		log.Error(err, "Marshal Failed")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(reply)
}

func handleSetAnswers(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleSetAnswers %s", request.URL.String())

	vars := mux.Vars(request)
	surveyID := vars["surveyId"]

	decoder := json.NewDecoder(request.Body)
	var answers []models.Answer
	err := decoder.Decode(&answers)
	if err != nil {
		log.Error(err, "Unmarshal Failed")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.IncrementSubmitCount(surveyID)
	if err != nil {
		log.Error(err, "Unable to Increment Result")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, answer := range answers {
		err = database.IncrementResult(surveyID, answer.QuestionID, answer.Answer)
		if err != nil {
			log.Error(err, "Unable to Increment Result")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	writer.WriteHeader(http.StatusNoContent)
}
