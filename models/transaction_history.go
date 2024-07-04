package models

import (
	"github.com/DarkPhoenix42/LibraLynx/db"
	"github.com/DarkPhoenix42/LibraLynx/types"
)

func GetUserTransactionHistory(user_id int) ([]types.ViewTransaction, error) {
	query := "SELECT t.*, b.title, u.username FROM transactions t JOIN books b ON t.book_id = b.book_id JOIN users u ON t.user_id = u.user_id WHERE t.user_id = ?"
	rows, err := db.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ParseViewTransactionRows(rows)
}

func GetOverallTransactionHistory() ([]types.ViewTransaction, error) {
	query := "SELECT t.*, b.title, u.username FROM transactions t JOIN books b ON t.book_id = b.book_id JOIN users u ON t.user_id = u.user_id"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ParseViewTransactionRows(rows)
}

func GetUserBorrowedTransactions(user_id int) ([]types.ViewTransaction, error) {
	query := "SELECT t.*, b.title, u.username FROM transactions t JOIN books b ON t.book_id = b.book_id JOIN users u ON t.user_id = u.user_id WHERE t.user_id = ? AND t.type = 'borrow' AND t.status = 'accepted'"
	rows, err := db.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ParseViewTransactionRows(rows)
}
