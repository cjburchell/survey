package main

import (
	"fmt"
	"github.com/cjburchell/settings-go"
	"github.com/cjburchell/tools-go/env"
	"net/http"
	"time"

	"github.com/cjburchell/survey/database"
	"github.com/cjburchell/survey/routes"
	"github.com/cjburchell/uatu-go"
	"github.com/gorilla/mux"
)

func main() {
	configFile := settings.Get(env.Get("SettingsFile", ""))
	logger := log.Create(configFile)

	logger.Print("Starting survey service")
	err := database.Connect(logger, configFile)
	if err != nil {
		logger.Fatalf(err, "Unable to connect to database")
	}

	logger.Print("Database Connected")

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "survey-ui/dist/survey-ui/index.html")
	})

	routes.Setup(router, logger)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("survey-ui/dist/survey-ui"))))

	logger.Print("Starting HTTP Server")
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
