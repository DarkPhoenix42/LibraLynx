package api

import (
	"log"
	"net/http"

	database "github.com/DarkPhoenix42/LibraLynx/db"
	admin_controllers "github.com/DarkPhoenix42/LibraLynx/pkg/controllers/admin"
	auth_controllers "github.com/DarkPhoenix42/LibraLynx/pkg/controllers/auth"
	user_controllers "github.com/DarkPhoenix42/LibraLynx/pkg/controllers/user"
	"github.com/DarkPhoenix42/LibraLynx/pkg/middleware"
	"github.com/DarkPhoenix42/LibraLynx/pkg/views"
	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.\n")
	}

	err = database.Connect()

	if err == nil {
		log.Printf("Successfully connected to the database.\n")
	} else {
		log.Fatalf("Error %s connecting to database.\n", err)
	}

	views.InitTemplates()

	http.HandleFunc("GET /register", auth_controllers.RegisterPage)
	http.HandleFunc("POST /register", auth_controllers.Register)

	http.HandleFunc("GET /login", auth_controllers.LoginPage)
	http.HandleFunc("POST /login", auth_controllers.Login)

	http.HandleFunc("GET /logout", auth_controllers.Logout)

	handleJWT := middleware.JwtMiddleware
	handleAdmin := middleware.AdminMiddleware

	userHomePage := http.HandlerFunc(user_controllers.UserHomePage)
	http.Handle("GET /", handleJWT(userHomePage))

	viewBooksPage := http.HandlerFunc(user_controllers.ViewBooksPage)
	http.Handle("GET /view_books", handleJWT(viewBooksPage))

	requestBorrowal := http.HandlerFunc(user_controllers.RequestBorrowal)
	http.Handle("POST /borrow_book/{book_id}", handleJWT(requestBorrowal))

	returnBookPage := http.HandlerFunc(user_controllers.ReturnBookPage)
	http.Handle("GET /return_book", handleJWT(returnBookPage))

	requestReturn := http.HandlerFunc(user_controllers.RequestReturn)
	http.Handle("POST /return_book/{transaction_id}", handleJWT(requestReturn))

	requestAdminPage := http.HandlerFunc(user_controllers.RequestAdminPage)
	http.Handle("GET /request_admin", handleJWT(requestAdminPage))

	requestAdmin := http.HandlerFunc(user_controllers.RequestAdmin)
	http.Handle("POST /request_admin", handleJWT(requestAdmin))

	transactionHistoryPage := http.HandlerFunc(user_controllers.TransactionHistoryPage)
	http.Handle("GET /transaction_history", handleJWT(transactionHistoryPage))

	// Admin routes
	adminHomePage := http.HandlerFunc(admin_controllers.AdminHomePage)
	http.Handle("GET /admin", handleJWT(handleAdmin(adminHomePage)))

	adminViewBooksPage := http.HandlerFunc(admin_controllers.ViewBooksPage)
	http.Handle("GET /admin/view_books", handleJWT(handleAdmin(adminViewBooksPage)))

	getBook := http.HandlerFunc(admin_controllers.GetBook)
	http.Handle("GET /admin/get_book/{book_id}", handleJWT(handleAdmin(getBook)))

	addBookPage := http.HandlerFunc(admin_controllers.AddBookPage)
	http.Handle("GET /admin/add_book", handleJWT(handleAdmin(addBookPage)))

	addBook := http.HandlerFunc(admin_controllers.AddBook)
	http.Handle("POST /admin/add_book", handleJWT(handleAdmin(addBook)))

	deleteBookPage := http.HandlerFunc(admin_controllers.DeleteBookPage)
	http.Handle("GET /admin/delete_book", handleJWT(handleAdmin(deleteBookPage)))

	deleteBook := http.HandlerFunc(admin_controllers.DeleteBook)
	http.Handle("POST /admin/delete_book", handleJWT(handleAdmin(deleteBook)))

	updateBookPage := http.HandlerFunc(admin_controllers.UpdateBookPage)
	http.Handle("GET /admin/update_book", handleJWT(handleAdmin(updateBookPage)))

	updateBook := http.HandlerFunc(admin_controllers.UpdateBook)
	http.Handle("POST /admin/update_book", handleJWT(handleAdmin(updateBook)))

	adminRequestsPage := http.HandlerFunc(admin_controllers.AdminRequestsPage)
	http.Handle("GET /admin/admin_requests", handleJWT(handleAdmin(adminRequestsPage)))

	acceptAdminRequest := http.HandlerFunc(admin_controllers.AcceptAdminRequest)
	http.Handle("POST /admin/accept_admin/{user_id}", handleJWT(handleAdmin(acceptAdminRequest)))

	rejectAdminRequest := http.HandlerFunc(admin_controllers.RejectAdminRequest)
	http.Handle("POST /admin/reject_admin/{user_id}", handleJWT(handleAdmin(rejectAdminRequest)))

	borrowRequestsPage := http.HandlerFunc(admin_controllers.BorrowRequestsPage)
	http.Handle("GET /admin/borrow_requests", handleJWT(handleAdmin(borrowRequestsPage)))

	acceptBorrowal := http.HandlerFunc(admin_controllers.AcceptBorrowal)
	http.Handle("POST /admin/accept_borrowal/{transaction_id}", handleJWT(handleAdmin(acceptBorrowal)))

	rejectBorrowal := http.HandlerFunc(admin_controllers.RejectBorrowal)
	http.Handle("POST /admin/reject_borrowal/{transaction_id}", handleJWT(handleAdmin(rejectBorrowal)))

	returnRequestsPage := http.HandlerFunc(admin_controllers.ReturnRequestsPage)
	http.Handle("GET /admin/return_requests", handleJWT(handleAdmin(returnRequestsPage)))

	acceptReturn := http.HandlerFunc(admin_controllers.AcceptReturn)
	http.Handle("POST /admin/accept_return/{transaction_id}", handleJWT(handleAdmin(acceptReturn)))

	rejectReturn := http.HandlerFunc(admin_controllers.RejectReturn)
	http.Handle("POST /admin/reject_return/{transaction_id}", handleJWT(handleAdmin(rejectReturn)))

	adminTransactionHistoryPage := http.HandlerFunc(admin_controllers.TransactionHistoryPage)
	http.Handle("GET /admin/transaction_history", handleJWT(handleAdmin(adminTransactionHistoryPage)))

	log.Printf("Server started at http://localhost:8080\n")
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error %s starting server.\n", err)
	}
}
