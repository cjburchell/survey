package routes

import (
	"encoding/json"
	"github.com/cjburchell/reefstatus-go/common/log"
	"github.com/cjburchell/survey/database"
	"github.com/cjburchell/survey/models"
	"github.com/gorilla/mux"
	"net/http"
)

func Setup(router *mux.Router) {
	surveyRoute := router.PathPrefix("/survey").Subrouter()
	surveyRoute.HandleFunc("/questions", handleGetQuestions).Methods("GET")
	surveyRoute.HandleFunc("/results", handleGetResults).Methods("GET")
	surveyRoute.HandleFunc("/results/{Id}", handleGetResultsForId).Methods("GET")
	surveyRoute.HandleFunc("/answers", handleSetAnswers).Methods("POST")

	surveyRoute.HandleFunc("/@status", func(w http.ResponseWriter, r *http.Request) {
		reply, _ := json.Marshal("Ok")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(reply)
	}).Methods("GET")
}

func handleGetQuestions(w http.ResponseWriter, r *http.Request) {
	log.Debugf("handleGetQuestions %s", r.URL.String())

	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "questions.json")
}

func handleGetResults(w http.ResponseWriter, r *http.Request) {
	log.Debugf("handleGetResults %s", r.URL.String())

	results, err := database.GetAllResults()
	if err != nil {
		log.Error(err, "GetAllResults Failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, err := json.Marshal(results)
	if err != nil {
		log.Error(err, "Marshal Failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

func handleGetResultsForId(w http.ResponseWriter, r *http.Request) {
	log.Debugf("handleGetResultsForId %s", r.URL.String())

	vars := mux.Vars(r)
	questionId := vars["Id"]
	results, err := database.GetResults(questionId)
	if err != nil {
		log.Error(err, "GetAllResults Failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, err := json.Marshal(results)
	if err != nil {
		log.Error(err, "Marshal Failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

func handleSetAnswers(w http.ResponseWriter, r *http.Request) {
	log.Debugf("handleSetAnswers %s", r.URL.String())
	var body []byte
	r.Body.Read(body)
	var answers []models.SurveyAnswer
	err := json.Unmarshal(body, &answers)
	if err != nil {
		log.Error(err, "Unmarshal Failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, answer := range answers {
		err = database.IncrementResult(answer.QuestionId, answer.Answer)
		if err != nil {
			log.Error(err, "Unable to Increment Result")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
