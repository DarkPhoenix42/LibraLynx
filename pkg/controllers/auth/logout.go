package controllers

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	utils.DeleteCookie(w, "jwt")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
