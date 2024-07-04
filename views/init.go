package views

import (
	"html/template"
)

var (
	Templates = map[string]*template.Template{}
)

func InitTemplates() {

	Templates["login"] = template.Must(template.ParseFiles("templates/auth/login.html", "templates/partials/head.html", "templates/partials/alert_box.html"))
	Templates["register"] = template.Must(template.ParseFiles("templates/auth/register.html", "templates/partials/head.html", "templates/partials/alert_box.html"))

	Templates["userHome"] = template.Must(template.ParseFiles("templates/user/home.html", "templates/partials/head.html", "templates/partials/alert_box.html"))
	Templates["viewBooks"] = template.Must(template.ParseFiles("templates/user/view_books.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["returnBook"] = template.Must(template.ParseFiles("templates/user/return_book.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["transactionHistory"] = template.Must(template.ParseFiles("templates/user/transaction_history.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["requestAdmin"] = template.Must(template.ParseFiles("templates/user/request_admin.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))

	Templates["adminHome"] = template.Must(template.ParseFiles("templates/admin/home.html", "templates/partials/head.html", "templates/partials/alert_box.html"))
	Templates["adminViewBooks"] = template.Must(template.ParseFiles("templates/admin/view_books.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["addBook"] = template.Must(template.ParseFiles("templates/admin/add_book.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["updateBook"] = template.Must(template.ParseFiles("templates/admin/update_book.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["deleteBook"] = template.Must(template.ParseFiles("templates/admin/delete_book.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["borrowRequests"] = template.Must(template.ParseFiles("templates/admin/borrow_requests.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["returnRequests"] = template.Must(template.ParseFiles("templates/admin/return_requests.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["adminRequests"] = template.Must(template.ParseFiles("templates/admin/admin_requests.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
	Templates["adminTransactionHistory"] = template.Must(template.ParseFiles("templates/admin/transaction_history.html", "templates/partials/head.html", "templates/partials/nav.html", "templates/partials/alert_box.html"))
}
