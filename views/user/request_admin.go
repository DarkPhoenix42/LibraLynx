package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/types"
	"github.com/DarkPhoenix42/LibraLynx/views"
)

func RequestAdmin(w http.ResponseWriter, data types.MessageData) {
	tmpl := views.Templates["requestAdmin"]
	tmpl.Execute(w, data)
}
