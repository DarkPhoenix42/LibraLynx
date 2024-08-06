package controllers

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/pkg/models"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
	views "github.com/DarkPhoenix42/LibraLynx/pkg/views/auth"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	message, msg_type := utils.GetAndClearMessage(w, r)
	views.Login(w, types.MessageData{Message: message, MessageType: msg_type})
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		utils.SetMessage(w, "Please enter username and password!", "error")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	user, err := models.GetUserByUsername(username)
	if err != nil {
		utils.SetMessage(w, "Invalid username or password!", "error")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return

	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		utils.SetMessage(w, "Invalid username or password!", "error")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Create session
	token, err := utils.CreateToken(user.UserID, user.IsAdmin)

	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	utils.SetCookie(w, "jwt", token)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
