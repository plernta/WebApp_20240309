package main

import (
	"net/http"
	"welldone/hong/cool_web_app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// router is making choices base on the url which handler to call
	r := mux.NewRouter()

	// connect URL to the handler
	r.HandleFunc("/", handlers.IndexPage)
	r.HandleFunc("/users/{username}", handlers.UserPage)

	// start the server
	http.ListenAndServe(":3000", r)
	//:3000 is the port number

	//http.ListenAndServe(":3000/about", r)
}
