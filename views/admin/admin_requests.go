package views

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/types"
	"github.com/DarkPhoenix42/LibraLynx/views"
)

func AdminRequests(w http.ResponseWriter, data types.ViewAdminRequestsData) {
	tmpl := views.Templates["adminRequests"]
	tmpl.Execute(w, data)
}
