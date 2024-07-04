package middleware

import (
	"net/http"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		is_admin := r.Context().Value("is_admin").(bool)
		if !is_admin {
			http.Redirect(w, r, "/request_admin", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
