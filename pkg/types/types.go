package types

import "time"

type User struct {
	UserID             int    `json:"user_id"`
	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	IsAdmin            bool   `json:"is_admin"`
	AdminRequestStatus string `json:"admin_request_status"` // Enum: none, pending
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
	Date          *time.Time `json:"date"`
	Type          string     `json:"type"`   // Enum: borrow, return
	Status        string     `json:"status"` // Enum: pending, accepted, rejected
}

type AdminRequest struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type MessageData struct {
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type ViewTransaction struct {
	TransactionID int        `json:"transaction_id"`
	BookID        int        `json:"book_id"`
	BookTitle     string     `json:"book_title"`
	UserID        int        `json:"user_id"`
	Username      string     `json:"username"`
	Date          *time.Time `json:"date"`
	Type          string     `json:"type"`   // Enum: borrow, return
	Status        string     `json:"status"` // Enum: pending, accepted, rejected
}

type ViewBooksData struct {
	Books       []Book
	MessageData MessageData
}

type ViewTransactionsData struct {
	Transactions []ViewTransaction
	MessageData  MessageData
}

type ViewAdminRequestsData struct {
	Requests    []AdminRequest
	MessageData MessageData
}
