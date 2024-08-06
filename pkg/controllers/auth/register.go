package controllers

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/pkg/models"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
	views "github.com/DarkPhoenix42/LibraLynx/pkg/views/auth"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	message, msg_type := utils.GetAndClearMessage(w, r)
	views.Register(w, types.MessageData{Message: message, MessageType: msg_type})
}

func Register(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm_password := r.FormValue("confirm_password")

	if username == "" || email == "" || password == "" || confirm_password == "" {
		utils.SetMessage(w, "Please fill in all fields!", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return

	}

	if password != confirm_password {
		utils.SetMessage(w, "Passwords do not match!", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	if !utils.CheckUsername(username) {
		utils.SetMessage(w, "Username must be between 3 and 20 characters long and can only contain letters, numbers, and underscores.", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return

	}

	if !utils.CheckPassword(password) {
		utils.SetMessage(w, "Password must be between atleast 7 characters long and must contain at least one uppercase letter, one lowercase letter, one number, and one special character.", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	if !utils.CheckEmail(email) {
		utils.SetMessage(w, "Invalid email address!", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	is_admin, err := models.CheckFirstUser()

	if err != nil {
		is_admin = false
	}

	user, _ := models.GetUserByUsername(username)
	if user != nil {
		utils.SetMessage(w, "Username already exists!", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	err = models.AddUser(username, email, password, is_admin)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	user, err = models.GetUserByUsername(username)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	
	user_id := user.UserID
	token, err := utils.CreateToken(user_id, is_admin)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	utils.SetCookie(w, "jwt", token)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
