package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/types"
	"github.com/DarkPhoenix42/LibraLynx/views"
)

func ViewBooks(w http.ResponseWriter, data types.ViewBooksData) {
	tmpl := views.Templates["adminViewBooks"]
	tmpl.Execute(w, data)
}

func AddBook(w http.ResponseWriter, data types.MessageData) {

	tmpl := views.Templates["addBook"]
	tmpl.Execute(w, data)
}
func UpdateBook(w http.ResponseWriter, data types.ViewBooksData) {

	tmpl := views.Templates["updateBook"]
	tmpl.Execute(w, data)
}

func DeleteBook(w http.ResponseWriter, data types.ViewBooksData) {
	tmpl := views.Templates["deleteBook"]
	tmpl.Execute(w, data)
}
