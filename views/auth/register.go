package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/types"
	"github.com/DarkPhoenix42/LibraLynx/views"
)

func Register(w http.ResponseWriter, data types.MessageData) {
	tmpl := views.Templates["register"]
	tmpl.Execute(w, data)
}
