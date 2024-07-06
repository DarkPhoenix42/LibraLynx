package controllers

import (
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/pkg/models"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
	views "github.com/DarkPhoenix42/LibraLynx/pkg/views/user"
)

func RequestAdminPage(w http.ResponseWriter, r *http.Request) {
	message, msg_type := utils.GetAndClearMessage(w, r)
	views.RequestAdmin(w, types.MessageData{Message: message, MessageType: msg_type})
}

func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	err := models.UpdateUserAdminRequestStatus(user_id, "pending")
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Admin request sent successfully!", "success")
	}

	http.Redirect(w, r, "/request_admin", http.StatusSeeOther)
}
