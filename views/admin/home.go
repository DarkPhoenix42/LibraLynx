package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/views"
)

func AdminHome(w http.ResponseWriter) {
	tmpl := views.Templates["adminHome"]
	tmpl.Execute(w, nil)
}
