package middleware

import (
	"fmt"
	"net/http"

	"github.com/awaisniaz/todo/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		fmt.Println("User-Agent:", userAgent)
		token := r.Header.Get("Authorization")
		fmt.Println("User-Agent:", token)
		valid, err := utils.VerifyToken(token)
		if valid == true {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusUnauthorized)

	})
}
