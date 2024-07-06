package controllers

import (
	"net/http"
	"strconv"

	"github.com/DarkPhoenix42/LibraLynx/pkg/models"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
	views "github.com/DarkPhoenix42/LibraLynx/pkg/views/admin"
)

func AdminRequestsPage(w http.ResponseWriter, r *http.Request) {
	requests, err := models.GetAdminRequests()
	message, msg_type := utils.GetAndClearMessage(w, r)
	if err != nil {
		views.AdminRequests(w, types.ViewAdminRequestsData{
			Requests:    []types.AdminRequest{},
			MessageData: types.MessageData{Message: "Internal server error!", MessageType: "error"}})
	} else {
		views.AdminRequests(w, types.ViewAdminRequestsData{
			Requests:    requests,
			MessageData: types.MessageData{Message: message, MessageType: msg_type}})
	}
}

func AcceptAdminRequest(w http.ResponseWriter, r *http.Request) {
	user_id, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid user ID!", "error")
		http.Redirect(w, r, "/admin/admin_requests", http.StatusSeeOther)
		return
	}
	err = models.UpdateUserAdminRequestStatus(user_id, "none")
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
		http.Redirect(w, r, "/admin/admin_requests", http.StatusSeeOther)
		return
	}
	err = models.MakeUserAdmin(user_id)
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Admin request accepted!", "success")
	}

	http.Redirect(w, r, "/admin/admin_requests", http.StatusSeeOther)
}

func RejectAdminRequest(w http.ResponseWriter, r *http.Request) {
	user_id, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		utils.SetMessage(w, "Invalid user ID!", "error")
		http.Redirect(w, r, "/admin/admin_requests", http.StatusSeeOther)
		return
	}

	err = models.UpdateUserAdminRequestStatus(user_id, "none")
	if err != nil {
		utils.SetMessage(w, "Internal server error!", "error")
	} else {
		utils.SetMessage(w, "Admin request rejected!", "success")
	}
	http.Redirect(w, r, "/admin/admin_requests", http.StatusSeeOther)

}
