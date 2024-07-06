package models

import "github.com/DarkPhoenix42/LibraLynx/db"

func AddBook(title, author, genre string, available_copies int) error {
	query := "INSERT INTO books (title, author, genre, available_copies) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, title, author, genre, available_copies)
	return err
}

func AddCopies(book_id, copies int) error {
	query := "UPDATE books SET available_copies = available_copies + ? WHERE book_id = ?"
	_, err := db.DB.Exec(query, copies, book_id)
	return err
}

func UpdateBook(book_id int, title, author, genre string, available_copies int) error {
	query := "UPDATE books SET title = ?, author = ?, genre = ?, available_copies = ? WHERE book_id = ?"
	_, err := db.DB.Exec(query, title, author, genre, available_copies, book_id)
	return err
}

func DeleteBook(book_id int) error {
	query := "DELETE FROM books WHERE book_id = ?"
	_, err := db.DB.Exec(query, book_id)
	return err
}

func IncrementAvailableCopies(book_id int) error {
	query := "UPDATE books SET available_copies = available_copies + 1 WHERE book_id = ?"
	_, err := db.DB.Exec(query, book_id)
	return err
}

func DecrementAvailableCopies(book_id int) error {
	query := "UPDATE books SET available_copies = available_copies - 1 WHERE book_id = ?"
	_, err := db.DB.Exec(query, book_id)
	return err
}
