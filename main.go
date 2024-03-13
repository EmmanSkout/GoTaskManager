package main

import (
	"log"
	"net/http"

	service "github.com/EmmanSkout/TaskManager/services"
	mux "github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files from the "docs" directory
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	r.HandleFunc("/tasks/add", service.HandleAdd).Methods("POST")
	r.HandleFunc("/tasks/modify", service.HandleModify).Methods("POST")
	r.HandleFunc("/tasks/load", service.HandleLoad).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./docs")))
	log.Fatal(http.ListenAndServe(":3000", r))
}
