package models

import (
	"database/sql"

	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
)

func ParseTransactionRows(rows *sql.Rows) ([]types.Transaction, error) {
	var transactions []types.Transaction
	for rows.Next() {
		var transaction types.Transaction
		err := rows.Scan(
			&transaction.TransactionID,
			&transaction.BookID,
			&transaction.UserID,
			&transaction.Date,
			&transaction.Type,
			&transaction.Status,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func ParseViewTransactionRows(rows *sql.Rows) ([]types.ViewTransaction, error) {
	var transactions []types.ViewTransaction
	for rows.Next() {
		var transaction types.ViewTransaction
		err := rows.Scan(
			&transaction.TransactionID,
			&transaction.BookID,
			&transaction.UserID,
			&transaction.Date,
			&transaction.Type,
			&transaction.Status,
			&transaction.BookTitle,
			&transaction.Username)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func ParseBookRows(rows *sql.Rows) ([]types.Book, error) {
	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(
			&book.BookID,
			&book.Title,
			&book.Author,
			&book.Genre,
			&book.AvailableCopies,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
