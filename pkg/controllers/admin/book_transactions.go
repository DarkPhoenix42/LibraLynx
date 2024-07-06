package controllers

import (
	"net/http"
	"strconv"

	"github.com/DarkPhoenix42/LibraLynx/pkg/models"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
	views "github.com/DarkPhoenix42/LibraLynx/pkg/views/admin"
)

func BorrowRequestsPage(w http.ResponseWriter, r *http.Request) {
	requests, err := models.GetPendingBorrowTransactions()
	message, msg_type := utils.GetAndClearMessage(w, r)
	if err != nil {
		views.BorrowRequests(w, types.ViewTransactionsData{
			Transactions: []types.ViewTransaction{},
			MessageData:  types.MessageData{Message: "Internal server error!", MessageType: "error"}})
	} else {
		views.BorrowRequests(w, types.ViewTransactionsData{
			Transactions: requests,
			MessageData:  types.MessageData{Message: message, MessageType: msg_type}})
	}
}

func ReturnRequestsPage(w http.ResponseWriter, r *http.Request) {
	requests, err := models.GetPendingReturnTransactions()
	message, msg_type := utils.GetAndClearMessage(w, r)
	if err != nil {
		views.ReturnRequests(w, types.ViewTransactionsData{
			Transactions: []types.ViewTransaction{},
			MessageData:  types.MessageData{Message: "Internal server error!", MessageType: "error"}})
	} else {
		views.ReturnRequests(w, types.ViewTransactionsData{
			Transactions: requests,
			MessageData:  types.MessageData{Message: message, MessageType: msg_type}})

	}
}

func TransactionHistoryPage(w http.ResponseWriter, r *http.Request) {
	transactions, err := models.GetOverallTransactionHistory()
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

func AcceptBorrowal(w http.ResponseWriter, r *http.Request) {
	transaction_id, err := strconv.Atoi(r.PathValue("transaction_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid transaction ID!", "error")
		http.Redirect(w, r, "/admin/borrow_requests", http.StatusSeeOther)
		return
	}
	err = models.UpdateTransactionStatus(transaction_id, "accepted")
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/admin/borrow_requests", http.StatusSeeOther)
		return
	}
	err = models.UpdateBorrowDate(transaction_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/admin/borrow_requests", http.StatusSeeOther)
		return
	}
	transaction, err := models.GetTransactionByID(transaction_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/admin/borrow_requests", http.StatusSeeOther)
		return
	}
	err = models.DecrementAvailableCopies(transaction.BookID)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Borrow request accepted!", "success")
	}
	http.Redirect(w, r, "/admin/borrow_requests", http.StatusSeeOther)
}

func AcceptReturn(w http.ResponseWriter, r *http.Request) {
	transaction_id, err := strconv.Atoi(r.PathValue("transaction_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid transaction ID!", "error")
		http.Redirect(w, r, "/admin/return_requests", http.StatusSeeOther)
		return
	}
	transaction, err := models.GetTransactionByID(transaction_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/admin/return_requests", http.StatusSeeOther)
		return
	}
	err = models.UpdateTransactionStatus(transaction_id, "accepted")
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/admin/return_requests", http.StatusSeeOther)
		return
	}
	err = models.UpdateReturnDate(transaction_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/admin/return_requests", http.StatusSeeOther)
		return
	}
	err = models.IncrementAvailableCopies(transaction.BookID)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Return request accepted!", "success")
	}
	http.Redirect(w, r, "/admin/return_requests", http.StatusSeeOther)
}

func RejectBorrowal(w http.ResponseWriter, r *http.Request) {
	transaction_id, err := strconv.Atoi(r.PathValue("transaction_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid transaction ID!", "error")
		http.Redirect(w, r, "/admin/borrow_requests", http.StatusSeeOther)
		return
	}
	err = models.UpdateTransactionStatus(transaction_id, "rejected")
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Borrow request rejected!", "success")
	}
	http.Redirect(w, r, "/admin/borrow_requests", http.StatusSeeOther)
}

func RejectReturn(w http.ResponseWriter, r *http.Request) {
	transaction_id, err := strconv.Atoi(r.PathValue("transaction_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid transaction ID!", "error")
		http.Redirect(w, r, "/admin/return_requests", http.StatusSeeOther)
		return
	}
	err = models.UpdateTransactionStatus(transaction_id, "rejected")
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Return request rejected!", "success")
	}
	http.Redirect(w, r, "/admin/return_requests", http.StatusSeeOther)
}
