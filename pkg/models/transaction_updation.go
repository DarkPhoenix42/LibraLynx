package models

import (
	"github.com/DarkPhoenix42/LibraLynx/db"
)

func InitiateBorrowTransaction(book_id, user_id int) error {
	query := "INSERT INTO transactions (book_id, user_id, type, status) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, book_id, user_id, "borrow", "pending")
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

func UpdateTransactionType(transaction_id int, transaction_type string) error {
	query := "UPDATE transactions SET type = ? WHERE transaction_id = ?"
	_, err := db.DB.Exec(query, transaction_type, transaction_id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateBorrowDate(transaction_id int) error {
	query := "UPDATE transactions SET borrow_date = NOW() WHERE transaction_id = ?"
	_, err := db.DB.Exec(query, transaction_id)

	if err != nil {
		return err
	}

	query = "UPDATE transactions SET due_date = DATE_ADD(borrow_date, INTERVAL 14 DAY) WHERE transaction_id = ?"
	_, err = db.DB.Exec(query, transaction_id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateReturnDate(transaction_id int) error {
	query := "UPDATE transactions SET return_date = NOW() WHERE transaction_id = ?"
	_, err := db.DB.Exec(query, transaction_id)
	if err != nil {
		return err
	}

	return nil
}
