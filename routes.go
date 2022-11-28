package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/create", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask_ID).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", router))

	return router

}
