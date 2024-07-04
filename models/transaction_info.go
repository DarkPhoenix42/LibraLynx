package models

import (
	"github.com/DarkPhoenix42/LibraLynx/db"
	"github.com/DarkPhoenix42/LibraLynx/types"
)

func GetTransactionByID(transaction_id int) (*types.Transaction, error) {
	var transaction types.Transaction
	query := "SELECT * FROM transactions WHERE transaction_id = ?"

	row := db.DB.QueryRow(query, transaction_id)
	err := row.Scan(
		&transaction.TransactionID,
		&transaction.BookID,
		&transaction.UserID,
		&transaction.BorrowDate,
		&transaction.DueDate,
		&transaction.ReturnDate,
		&transaction.Fine,
		&transaction.Type,
		&transaction.Status)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func GetBookTransactions(book_id int) ([]types.Transaction, error) {
	query := "SELECT * FROM transactions WHERE book_id = ?"
	rows, err := db.DB.Query(query, book_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ParseTransactionRows(rows)
}

func GetUserBookTransactions(user_id, book_id int) ([]types.Transaction, error) {
	query := "SELECT * FROM transactions WHERE user_id = ? AND book_id = ?"
	rows, err := db.DB.Query(query, user_id, book_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ParseTransactionRows(rows)
}
