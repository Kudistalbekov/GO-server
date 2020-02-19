package main

import (
	"log"
	"net/http"
	"projects/server/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() //creating new multiplexer
	router.HandleFunc("/create", handlers.Reg).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
