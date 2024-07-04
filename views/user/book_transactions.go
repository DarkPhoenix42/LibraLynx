package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/types"
	"github.com/DarkPhoenix42/LibraLynx/views"
)

func ViewBooks(w http.ResponseWriter, data types.ViewBooksData) {
	tmpl := views.Templates["viewBooks"]
	tmpl.Execute(w, data)
}

func ReturnBook(w http.ResponseWriter, data types.ViewTransactionsData) {
	tmpl := views.Templates["returnBook"]
	tmpl.Execute(w, data)
}

func TransactionHistory(w http.ResponseWriter, data types.ViewTransactionsData) {
	tmpl := views.Templates["transactionHistory"]

	tmpl.Execute(w, data)
}
