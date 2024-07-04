package models

import (
	"github.com/DarkPhoenix42/LibraLynx/db"
	"github.com/DarkPhoenix42/LibraLynx/types"
)

func GetPendingBorrowTransactions() ([]types.ViewTransaction, error) {
	query := "SELECT t.*, b.title, u.username FROM transactions t JOIN books b ON t.book_id = b.book_id JOIN users u ON t.user_id = u.user_id WHERE t.type = 'borrow' AND t.status = 'pending'"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ParseViewTransactionRows(rows)
}

func GetPendingReturnTransactions() ([]types.ViewTransaction, error) {
	query := "SELECT t.*, b.title, u.username FROM transactions t JOIN books b ON t.book_id = b.book_id JOIN users u ON t.user_id = u.user_id WHERE t.type = 'return' AND t.status = 'pending'"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ParseViewTransactionRows(rows)
}
