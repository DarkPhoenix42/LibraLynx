package models

import (
	"github.com/DarkPhoenix42/LibraLynx/db"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
)

func GetBookByID(book_id int) (*types.Book, error) {
	query := "SELECT * FROM books WHERE book_id = ?"
	row := db.DB.QueryRow(query, book_id)

	var book types.Book
	err := row.Scan(
		&book.BookID,
		&book.Title,
		&book.Author,
		&book.Genre,
		&book.AvailableCopies)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func GetBookByInfo(title, author, genre string) (*types.Book, error) {
	query := "SELECT * FROM books WHERE title = ? AND author = ? AND genre = ?"
	row := db.DB.QueryRow(query, title, author, genre)

	var book types.Book
	err := row.Scan(
		&book.BookID,
		&book.Title,
		&book.Author,
		&book.Genre,
		&book.AvailableCopies)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func GetAllBooks() ([]types.Book, error) {
	query := "SELECT * FROM books"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ParseBookRows(rows)
}
