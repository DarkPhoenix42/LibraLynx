package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/types"
	"github.com/DarkPhoenix42/LibraLynx/views"
)

func BorrowRequests(w http.ResponseWriter, data types.ViewTransactionsData) {
	tmpl := views.Templates["borrowRequests"]
	tmpl.Execute(w, data)
}

func ReturnRequests(w http.ResponseWriter, data types.ViewTransactionsData) {
	tmpl := views.Templates["returnRequests"]
	tmpl.Execute(w, data)
}

func TransactionHistory(w http.ResponseWriter, data types.ViewTransactionsData) {
	tmpl := views.Templates["adminTransactionHistory"]
	tmpl.Execute(w, data)
}
