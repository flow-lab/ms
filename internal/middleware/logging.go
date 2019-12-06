package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging logs all requests with its path and the processing time
func Logging(log *log.Logger) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
