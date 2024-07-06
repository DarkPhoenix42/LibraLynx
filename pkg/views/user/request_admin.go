package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/views"
)

func RequestAdmin(w http.ResponseWriter, data types.MessageData) {
	tmpl := views.Templates["requestAdmin"]
	tmpl.Execute(w, data)
}
