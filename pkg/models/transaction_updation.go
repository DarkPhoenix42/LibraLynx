package models

import (
	"github.com/DarkPhoenix42/LibraLynx/db"
)

func InitiateBorrowTransaction(book_id, user_id int) error {
	query := "INSERT INTO transactions (book_id, user_id, type, status) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, book_id, user_id, "borrow", "pending")
	return err
}

func InitiateReturnTransaction(book_id, user_id int) error {
	query := "INSERT INTO transactions (book_id, user_id, type, status) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, book_id, user_id, "return", "pending")
	return err
}

func UpdateTransactionStatus(transaction_id int, status string) error {
	query := "UPDATE transactions SET status = ? WHERE transaction_id = ?"
	_, err := db.DB.Exec(query, status, transaction_id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTransactionDate(transaction_id int) error {
	query := "UPDATE transactions SET date = NOW() WHERE transaction_id = ?"
	_, err := db.DB.Exec(query, transaction_id)
	if err != nil {
		return err
	}

	return nil
}
