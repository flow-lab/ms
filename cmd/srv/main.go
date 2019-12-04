package main

import (
	"fmt"
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
	http.HandleFunc("/health", Chain(Health, OnlyMethod("GET"), Logging()))
	return http.ListenAndServe(":80", nil)
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

func Health(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Ok")
}
