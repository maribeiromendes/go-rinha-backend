package middlewares

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s - %s",r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
