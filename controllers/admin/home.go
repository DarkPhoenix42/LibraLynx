package controllers

import (
	"net/http"

	views "github.com/DarkPhoenix42/LibraLynx/views/admin"
)

func AdminHomePage(w http.ResponseWriter, r *http.Request) {
	views.AdminHome(w)
}
