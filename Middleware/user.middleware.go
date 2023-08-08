package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/awaisniaz/todo/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		valid, err := utils.VerifyToken(token)
		fmt.Println(valid)
		if err == nil {
			ctx := context.WithValue(r.Context(), "userID", valid.UserID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusUnauthorized)

	})
}
