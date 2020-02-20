package main

import (
	"log"
	"net/http"
	"projects/server/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter() //creating new multiplexer
	router.HandleFunc("/user/", handlers.RegPost).Methods("POST")
	router.HandleFunc("/user/", handlers.RegGet).Methods("GET").Queries("email", "")
	log.Fatal(http.ListenAndServe(":8080", router))
}
