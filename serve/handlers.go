package main

import (
	"html/template"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

var funcMap = template.FuncMap{}

var tmpl, _ = template.New("main").
	Funcs(funcMap).
	ParseGlob("templates/*.gohtml")

func renderHTML(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {

	err := tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		http.ServeFile(w, r, "public/50x.html")
		log.WithField("uri", r.URL).
			Error(err)
		return
	}

	log.WithField("uri", r.URL).Debug("Rendered page.")
}

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/404.html")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	renderHTML(w, r, "home", data)
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	renderHTML(w, r, "mathematics", data)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	renderHTML(w, r, "datascience", data)
}

func musicHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	renderHTML(w, r, "music", data)
}
