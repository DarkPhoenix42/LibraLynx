package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/pkg/views"
)

func AdminHome(w http.ResponseWriter) {
	tmpl := views.Templates["adminHome"]
	tmpl.Execute(w, nil)
}
