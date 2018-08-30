package models

// Book struct
type Book struct {
	ID 		string 	`json:"id"`
	Isbn	string 	`json:"isbn"`
	Title	string 	`json:"title"`
	Author	*Author	`json:"author"`
}

// Author struct
type Author struct {
	FirstName	string 	`json:"firstname"`
	LastName	string	`json:"lastname"`
}

var Books []Book

// Return the book struct
func BookModel() []Book {
	Books = []Book{
		{ID:"1", Isbn:"1201", Title:"A Faith in Fate", Author: &Author{ FirstName: "Kenechukwu", LastName: "Nnamani" }},
		{ID:"2", Isbn:"7820", Title:"Go in the Fruit", Author: &Author{ FirstName: "John", LastName: "Doe" }},
	}

	return Books
}