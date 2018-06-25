package routes

import (
	"encoding/json"
	"github.com/cjburchell/reefstatus-go/common/log"
	"github.com/cjburchell/survey/database"
	"github.com/cjburchell/survey/models"
	"github.com/gorilla/mux"
	"net/http"
)

// Sets up the routes
func Setup(router *mux.Router) {
	surveyRoute := router.PathPrefix("/survey").Subrouter()
	surveyRoute.HandleFunc("/{surveyId}", handleGetSurvey).Methods("GET")
	surveyRoute.HandleFunc("/{surveyId}/results", handleGetResults).Methods("GET")
	surveyRoute.HandleFunc("/{surveyId}/results/{questionId}", handleGetResultsForQuestion).Methods("GET")
	surveyRoute.HandleFunc("/{surveyId}/answers", handleSetAnswers).Methods("POST")
	surveyRoute.HandleFunc("/{surveyId}/count", handleGetSurveyCount).Methods("GET")

	router.HandleFunc("/@status", func(writer http.ResponseWriter, _ *http.Request) {
		reply, _ := json.Marshal("Ok")
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(reply)
	}).Methods("GET")
}

func handleGetSurvey(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleGetQuestions %s", request.URL.String())

	vars := mux.Vars(request)
	surveyId := vars["surveyId"]

	survey, err := database.GetSurvey(surveyId)
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

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(reply)
}

func handleGetSurveyCount(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleGetSurveyCount %s", request.URL.String())

	vars := mux.Vars(request)
	surveyId := vars["surveyId"]

	count := database.GetSubmitCount(surveyId)

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

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(reply)
}

func handleGetResults(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleGetResults %s", request.URL.String())

	vars := mux.Vars(request)
	surveyId := vars["surveyId"]

	results, err := database.GetAllResults(surveyId)
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

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(reply)
}

func handleGetResultsForQuestion(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleGetResultsForId %s", request.URL.String())

	vars := mux.Vars(request)
	questionId := vars["questionId"]
	surveyId := vars["surveyId"]

	results, err := database.GetResults(surveyId, questionId)
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

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(reply)
}

func handleSetAnswers(writer http.ResponseWriter, request *http.Request) {
	log.Debugf("handleSetAnswers %s", request.URL.String())

	vars := mux.Vars(request)
	surveyId := vars["surveyId"]

	decoder := json.NewDecoder(request.Body)
	var answers []models.Answer
	err := decoder.Decode(&answers)
	if err != nil {
		log.Error(err, "Unmarshal Failed")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.IncrementSubmitCount(surveyId)
	if err != nil {
		log.Error(err, "Unable to Increment Result")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, answer := range answers {
		err = database.IncrementResult(surveyId, answer.QuestionId, answer.Answer)
		if err != nil {
			log.Error(err, "Unable to Increment Result")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	writer.WriteHeader(http.StatusNoContent)
}
