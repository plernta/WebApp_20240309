package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	//open index.html template and send its content to the client

	//find all templates in the template directory
	templates, err := filepath.Glob("templates/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//read all templates from the disk
	t, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//send template to the client
	err = t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func AboutPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Hong\n"))
}
