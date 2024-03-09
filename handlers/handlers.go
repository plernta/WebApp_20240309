package handlers

import (
	"net/http"

	"log" //print on terminal

	"github.com/gorilla/mux"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	//open index.html template and send its content to the client
	renderTemplate(w, "index.html", nil)
}

var users = map[string]map[string]any{
	"hong": {
		"Name": "Hong",
		"City": "Bangkok",
	},

	"john": {
		"Name": "John",
		"City": "Ohio",
	},

	"lisa": {
		"Name": "Lisa",
		"City": "Manhattan",
	},
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, ok := users[username]
	log.Printf("User: %v, ok: %v", user, ok)
	if !ok {
		http.NotFound(w, r)
		return
	}

	renderTemplate(w, "user.html", user)
}
