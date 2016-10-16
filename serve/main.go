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
	// This is looking for the asset files near the build binary.
	publicAssetsPath := "/Users/aford/go/src/github.com/gaford/gaustinford.com/serve/public/assets/"
	r.PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(publicAssetsPath)))).
		Methods("GET")
	r.HandleFunc("/heartbeat", heartbeatHandler).Methods("GET")

	serve := &http.Server{
		Addr:         ":8082",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Infof("Listening on port %d...", 8082)
	log.Fatal(serve.ListenAndServe())
}
