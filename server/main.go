package main

import (
	. "GoAPI/server/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// init Route
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/api/books", CreateBooks).Methods("POST")
	r.HandleFunc("/api/books/{id} ", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")

	// Listen and Serve
	log.Fatal(http.ListenAndServe(":8000", r))

}