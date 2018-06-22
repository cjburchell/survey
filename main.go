package main

import (
	"encoding/json"
	"fmt"
	"github.com/cjburchell/reefstatus-go/common/log"
	"github.com/cjburchell/survey/database"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to database %s", err.Error())
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "survey-ui/dist/survey-ui/index.html")
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("survey-ui/dist/survey-ui"))))

	r.HandleFunc("/questions", handleGetQuestions).Methods("GET")
	r.HandleFunc("/results", handleGetResults).Methods("GET")
	r.HandleFunc("/results/{Id}", handleGetResultsForId).Methods("GET")
	r.HandleFunc("/answers", handleSetAnswers).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8088",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf(err.Error())
	}
}

/*type Question struct {
	Id      string   `json:"id"`
	Text    string   `json:"text"`
	Type    string   `json:"type"`
	Choices []string `json:"choices"`
}*/

type SurveyAnswer struct {
	QuestionId string `json:"questionId"`
	Answer     string `json:"answer"`
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
	var answers []SurveyAnswer
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
