package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/heartbeat", heartbeatHandler).Methods("GET")

	serve := &http.Server{
		Addr:         ":8082",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
