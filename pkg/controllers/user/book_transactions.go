package controllers

import (
	"net/http"
	"strconv"

	"github.com/DarkPhoenix42/LibraLynx/pkg/models"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
	views "github.com/DarkPhoenix42/LibraLynx/pkg/views/user"
)

func ViewBooksPage(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	message, msg_type := utils.GetAndClearMessage(w, r)

	if err != nil {
		views.ViewBooks(w, types.ViewBooksData{
			Books:       []types.Book{},
			MessageData: types.MessageData{Message: "Internal server error!", MessageType: "error"}})
	} else {
		views.ViewBooks(w, types.ViewBooksData{
			Books:       books,
			MessageData: types.MessageData{Message: message, MessageType: msg_type}})
	}
}

func RequestBorrowal(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	book_id, err := strconv.Atoi(r.PathValue("book_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid book ID!", "error")
		http.Redirect(w, r, "/view_books", http.StatusSeeOther)
		return
	}

	book, err := models.GetBookByID(book_id)
	if err != nil || book.AvailableCopies == 0 {
		utils.SetMessage(w, "Book not available!", "error")
		http.Redirect(w, r, "/view_books", http.StatusSeeOther)
		return
	}

	user_book_transactions, err := models.GetUserBookTransactions(user_id, book_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/view_books", http.StatusSeeOther)
		return
	}

	for _, transaction := range user_book_transactions {
		if transaction.Status == "pending" && transaction.Type == "borrow" {
			utils.SetMessage(w, "You already have a pending borrow request for this book!", "error")
			http.Redirect(w, r, "/view_books", http.StatusSeeOther)
			return

		} else if transaction.Status == "pending" && transaction.Type == "return" {
			utils.SetMessage(w, "You already have a pending return request for this book!", "error")
			http.Redirect(w, r, "/view_books", http.StatusSeeOther)
			return

		} else if transaction.Status == "accepted" && transaction.Type == "borrow" {
			utils.SetMessage(w, "You already have this book!", "error")
			http.Redirect(w, r, "/view_books", http.StatusSeeOther)
			return
		}
	}
	
	err = models.InitiateBorrowTransaction(book_id, user_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Borrow request sent successfully!", "success")
	}

	http.Redirect(w, r, "/view_books", http.StatusSeeOther)

}

func ReturnBookPage(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	transactions, err := models.GetUserBorrowedTransactions(user_id)
	message, msg_type := utils.GetAndClearMessage(w, r)

	if err != nil {
		views.ReturnBook(w, types.ViewTransactionsData{
			Transactions: []types.ViewTransaction{},
			MessageData:  types.MessageData{Message: message, MessageType: msg_type}})
	} else {
		views.ReturnBook(w, types.ViewTransactionsData{
			Transactions: transactions,
			MessageData:  types.MessageData{Message: message, MessageType: msg_type}})
	}
}

func RequestReturn(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	transaction_id, err := strconv.Atoi(r.PathValue("transaction_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid transaction ID!", "error")
		http.Redirect(w, r, "/return_book", http.StatusSeeOther)
		return
	}

	transaction, err := models.GetTransactionByID(transaction_id)
	if err != nil {
		utils.SetMessage(w, "Invalid transaction ID!", "error")
		http.Redirect(w, r, "/return_book", http.StatusSeeOther)
		return
	}

	if transaction.UserID != user_id {
		utils.SetMessage(w, "Cannot return book that is not borrowed by you!", "error")
		http.Redirect(w, r, "/return_book", http.StatusSeeOther)
		return
	}

	err = models.InitiateReturnTransaction(transaction.BookID, user_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/return_book", http.StatusSeeOther)
		return
	}

	err = models.UpdateTransactionStatus(transaction_id, "archived")
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Return request sent successfully!", "success")
	}

	http.Redirect(w, r, "/return_book", http.StatusSeeOther)
}

func TransactionHistoryPage(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	transactions, err := models.GetUserTransactionHistory(user_id)
	message, msg_type := utils.GetAndClearMessage(w, r)

	if err != nil {
		views.TransactionHistory(w, types.ViewTransactionsData{
			Transactions: []types.ViewTransaction{},
			MessageData:  types.MessageData{Message: "Internal server error!", MessageType: "error"}})
	} else {
		views.TransactionHistory(w, types.ViewTransactionsData{
			Transactions: transactions,
			MessageData:  types.MessageData{Message: message, MessageType: msg_type}})
	}

}
