package controllers

import (
	"net/http"

	views "github.com/DarkPhoenix42/LibraLynx/views/user"
)

func UserHomePage(w http.ResponseWriter, r *http.Request) {
	is_admin := r.Context().Value("is_admin").(bool)
	if is_admin {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	views.UserHome(w)
}
