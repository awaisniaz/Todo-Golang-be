package middleware

import (
	"fmt"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		fmt.Println("User-Agent:", userAgent)
		token := r.Header.Get("Authorization")
		fmt.Println("User-Agent:", token)
		next.ServeHTTP(w, r)
	})
}
