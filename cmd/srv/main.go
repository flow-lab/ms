package main

import (
	database "github.com/flow-lab/ms/internal/platform"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "postgres"
	}

	user := os.Getenv("DB_USER")

	pass := os.Getenv("DB_PASSWORD")

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	disableTLS := false
	if os.Getenv("DB_DISABLE_TLS") == "true" {
		disableTLS = true
	}

	c := database.Config{
		Name:       name,
		User:       user,
		Password:   pass,
		Host:       host,
		DisableTLS: disableTLS,
	}

	log.Printf("listening on localhost:%s", port)
	http.HandleFunc("/health", Chain(Health(c), OnlyMethod("GET"), Logging()))
	return http.ListenAndServe(":"+port, nil)
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, ms ...Middleware) http.HandlerFunc {
	for _, m := range ms {
		f = m(f)
	}
	return f
}

// Logging logs all requests with its path and the processing time
func Logging() Middleware {
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

// OnlyMethod ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func OnlyMethod(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}
