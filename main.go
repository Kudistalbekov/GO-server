package main

import (
	"log"
	"net/http"
	"os"
	"projects/server/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	log.SetOutput(os.Stdout)
	router := mux.NewRouter() //creating new multiplexer
	router.HandleFunc("/user/", handlers.HandleError(handlers.RegPost)).Methods("POST")
	router.HandleFunc("/user/", handlers.HandleError(handlers.RegGet)).Methods("GET").Queries("email", "")
	router.HandleFunc("/user/", handlers.HandleError(handlers.ReqDelete)).Methods("DELETE").Queries("email", "")
	router.HandleFunc("/user/", handlers.HandleError(handlers.ReqPut)).Methods("PUT").Queries("email", "")
	log.Fatal(http.ListenAndServe(":8080", router))
}
