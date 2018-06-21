package main

import (
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/test/indeximport-issue72/index"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "survey-ui/dist/survey-ui/index.html")
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("survey-ui/dist/survey-ui"))))

	r.HandleFunc("/questions", handleGetQuestions).Methods("GET")
	r.HandleFunc("/answers", handleGetAnswers).Methods("GET")
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

func handleGetQuestions(w http.ResponseWriter, _ *http.Request) {
	reply, _ := json.Marshal("it works!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

func handleGetAnswers(w http.ResponseWriter, _ *http.Request) {
	reply, _ := json.Marshal("it works!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

func handleSetAnswers(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
