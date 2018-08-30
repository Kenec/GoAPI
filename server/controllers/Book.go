package controllers

import (
	"net/http"
	"encoding/json"
	. "GoAPI/server/models"
	"github.com/gorilla/mux"
	"strconv"
	"math/rand"
)

// ./GET Get Books controller
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(BookModel())
}

// ./GET Get Book controller
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop through the books
	for _, item := range BookModel() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

// ./POST Create Book controller
func CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}

// ./UPDATE Update Book controller
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

// ./DELETE Delete Book controller
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range BookModel() {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Books)
}