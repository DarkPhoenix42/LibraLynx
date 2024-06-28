package types

import "time"

type User struct {
	UserID             int    `json:"user_id"`
	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	IsAdmin            bool   `json:"is_admin"`
	AdminRequestStatus string `json:"admin_request_status"`
}

type Book struct {
	BookID          int    `json:"book_id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Genre           string `json:"genre"`
	AvailableCopies int    `json:"available_copies"`
}

type Transaction struct {
	TransactionID int        `json:"transaction_id"`
	BookID        int        `json:"book_id"`
	UserID        int        `json:"user_id"`
	BorrowDate    time.Time  `json:"borrow_date"`
	DueDate       time.Time  `json:"due_date"`
	ReturnDate    *time.Time `json:"return_date,omitempty"`
	Fine          float64    `json:"fine"`
	Type          string     `json:"type"`   // Enum: borrow, return
	Status        string     `json:"status"` // Enum: pending, accepted, rejected
}
