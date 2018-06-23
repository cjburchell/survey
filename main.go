package main

import (
	"fmt"
	"github.com/cjburchell/reefstatus-go/common/log"
	"github.com/cjburchell/survey/database"
	"github.com/cjburchell/survey/routes"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	log.Print("Starting survey service")
	err := database.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to database %s", err.Error())
	}

	log.Print("Database Connected")

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "survey-ui/dist/survey-ui/index.html")
	})

	routes.Setup(router)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("survey-ui/dist/survey-ui"))))

	log.Print("Starting HTTP Server")
	server := &http.Server{
		Handler:      router,
		Addr:         ":8088",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf(err.Error())
	}
}
