package middleware

import (
	"context"
	"net/http"

	"github.com/DarkPhoenix42/LibraLynx/utils"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		user_id, is_admin, err := utils.DecodeToken(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", user_id)
		ctx = context.WithValue(ctx, "is_admin", is_admin)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
