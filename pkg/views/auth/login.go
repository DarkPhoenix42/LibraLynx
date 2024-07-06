package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/views"
)

func Login(w http.ResponseWriter, data types.MessageData) {
	tmpl := views.Templates["login"]
	tmpl.Execute(w, data)
}
