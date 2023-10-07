package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "123456" {
			logrus.Errorf("unauthorized access attempt from %v", r.RemoteAddr)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
