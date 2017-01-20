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

	r.HandleFunc("/mathematics", mathHandler).Methods("GET")
	r.HandleFunc("/datascience", dataHandler).Methods("GET")
	r.HandleFunc("/music", musicHandler).Methods("GET")

	r.PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(Config.PublicAssetsPath)))).
		Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	serve := &http.Server{
		Addr:         ":" + Config.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Infof("Listening on port %s...", Config.Port)
	log.Fatal(serve.ListenAndServe())
}
