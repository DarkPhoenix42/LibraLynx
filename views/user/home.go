package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/views"
)

func UserHome(w http.ResponseWriter) {
	tmpl := views.Templates["userHome"]
	tmpl.Execute(w, nil)
}
