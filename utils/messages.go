package utils

import (
	"net/http"
)

func SetMessage(w http.ResponseWriter, message string, msg_type string) {
	SetCookie(w, "message", message)
	SetCookie(w, "msg_type", msg_type)
}

func GetAndClearMessage(w http.ResponseWriter, r *http.Request) (string, string) {
	message_cookie, _ := r.Cookie("message")
	msg_type_cookie, _ := r.Cookie("msg_type")

	if message_cookie == nil {
		return "", ""
	}

	ClearMessage(w)
	return message_cookie.Value, msg_type_cookie.Value
}

func ClearMessage(w http.ResponseWriter) {
	DeleteCookie(w, "message")
	DeleteCookie(w, "msg_type")
}
