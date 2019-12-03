package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
}

type Route struct {
	Logger  bool
	Handler http.Handler
}

type App struct {
	Health *Route
}

func run() error {
	app := &App{
		Health: &Route{
			Logger: true,
		},
	}

	fmt.Println("started ...")
	return http.ListenAndServe(":8080", app)
}

type Health struct{}

func (h Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	if len(head) == 0 {
		_, _ = w.Write([]byte("ok"))
	} else {
		http.Error(w, "not found", http.StatusNotFound)
	}
}

func shiftPath(p string) (string, string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	var head, tail string
	if i <= 0 {
		head = p[1:]
		tail = "/"
	} else {
		head = p[1:i]
		tail = p[i:]
	}
	return head, tail
}

func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var next *Route
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	if len(head) == 0 {
		next = &Route{
			Logger: true,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte("index"))
			}),
		}
	} else if head == "health" {
		var i interface{} = Health{}
		next = &Route{
			Logger:  true,
			Handler: i.(http.Handler),
		}
	} else {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if next.Logger {
		next.Handler = h.log(next.Handler)
	}

	next.Handler.ServeHTTP(w, r)

}

func (h *App) log(handler http.Handler) http.Handler {
	// TODO [grokrz]: add logging
	return handler
}
