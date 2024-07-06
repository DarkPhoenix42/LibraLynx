package controllers

import (
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/DarkPhoenix42/LibraLynx/pkg/models"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
	views "github.com/DarkPhoenix42/LibraLynx/pkg/views/admin"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	book_id, err := strconv.Atoi(r.PathValue("book_id"))
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		w.Write([]byte{})
		return
	}
	book, err := models.GetBookByID(book_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		w.Write([]byte{})
		return
	}

	bookJSON, err := json.Marshal(book)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		w.Write([]byte{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bookJSON)
}

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

func AddBookPage(w http.ResponseWriter, r *http.Request) {
	message, msg_type := utils.GetAndClearMessage(w, r)
	views.AddBook(w, types.MessageData{Message: message, MessageType: msg_type})
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	genre := r.FormValue("genre")
	available_copies, err := strconv.Atoi(r.FormValue("available_copies"))

	if err != nil {
		utils.SetMessage(w, "Invalid number of available copies!", "error")
		http.Redirect(w, r, "/admin/add_book", http.StatusSeeOther)
		return
	}

	book, _ := models.GetBookByInfo(title, author, genre)
	if book != nil {
		err = models.AddCopies(book.BookID, available_copies)
		if err != nil {
			utils.SetMessage(w, "Internal server error!", "error")
		} else {
			utils.SetMessage(w, "Book already exists, so added copies instead!", "success")
		}
	} else {

		err = models.AddBook(title, author, genre, available_copies)
		if err != nil {
			utils.SetMessage(w, "Internal server error!", "error")
		} else {
			utils.SetMessage(w, "Book added successfully!", "success")
		}
	}

	http.Redirect(w, r, "/admin/add_book", http.StatusSeeOther)
}

func UpdateBookPage(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	message, msg_type := utils.GetAndClearMessage(w, r)
	if err != nil {
		views.UpdateBook(w, types.ViewBooksData{
			Books:       []types.Book{},
			MessageData: types.MessageData{Message: "Internal server error!", MessageType: "error"}})
	} else {
		views.UpdateBook(w, types.ViewBooksData{
			Books:       books,
			MessageData: types.MessageData{Message: message, MessageType: msg_type}})
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book_id, err := strconv.Atoi(r.FormValue("book_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid book ID!", "error")
		http.Redirect(w, r, "/admin/update_book", http.StatusSeeOther)
		return
	}
	title := r.FormValue("title")
	author := r.FormValue("author")
	genre := r.FormValue("genre")
	available_copies, err := strconv.Atoi(r.FormValue("available_copies"))
	if err != nil {
		utils.SetMessage(w, "Invalid number of available copies!", "error")
		http.Redirect(w, r, "/admin/update_book", http.StatusSeeOther)
		return
	}
	err = models.UpdateBook(book_id, title, author, genre, available_copies)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Book updated successfully!", "success")

	}
	http.Redirect(w, r, "/admin/update_book", http.StatusSeeOther)
}

func DeleteBookPage(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	message, msg_type := utils.GetAndClearMessage(w, r)
	if err != nil {
		views.DeleteBook(w, types.ViewBooksData{
			Books:       []types.Book{},
			MessageData: types.MessageData{Message: "Internal server error!", MessageType: "error"}})
	} else {
		views.DeleteBook(w, types.ViewBooksData{
			Books:       books,
			MessageData: types.MessageData{Message: message, MessageType: msg_type}})

	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	book_id, err := strconv.Atoi(r.FormValue("book_id"))

	if err != nil {
		utils.SetMessage(w, "Invalid book ID!", "error")
		http.Redirect(w, r, "/admin/delete_book", http.StatusSeeOther)
		return
	}

	book_transactions, err := models.GetBookTransactions(book_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/admin/delete_book", http.StatusSeeOther)
		return
	}

	okay_to_delete := true
	for _, book_transaction := range book_transactions {
		if (book_transaction.Type == "borrow" && book_transaction.Status == "accepted") || (book_transaction.Status == "pending") {
			okay_to_delete = false
			break
		}
	}

	if !okay_to_delete {
		utils.SetMessage(w, "Cannot delete book with pending transactions!", "error")
		http.Redirect(w, r, "/admin/delete_book", http.StatusSeeOther)
		return
	}
	err = models.DeleteBook(book_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Book deleted successfully!", "success")
	}

	http.Redirect(w, r, "/admin/delete_book", http.StatusSeeOther)
}
