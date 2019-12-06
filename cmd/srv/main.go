package main

import (
	"context"
	"errors"
	"expvar"
	"fmt"
	"github.com/flow-lab/ms/internal/middleware"
	database "github.com/flow-lab/ms/internal/platform"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
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

	expvar.NewString("version").Set(version)
	expvar.NewString("commit").Set(commit)
	expvar.NewString("date").Set(date)

	l := log.New(os.Stdout, fmt.Sprintf("MS : (%s, %s) : ", version, short(commit)), log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// db
	l.Println("initializing database support")
	db, err := database.Open(c)
	if err != nil {
		return err
	}
	defer func() {
		log.Printf("closing db connection")
		_ = db.Close()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	l.Println("initializing api server")
	readTimeout, _ := time.ParseDuration("5s")
	writeTimeout, _ := time.ParseDuration("5s")
	apiSrv := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	defer l.Printf("completed")
	go func(log *log.Logger) {
		log.Printf("api server listening on %s", apiSrv.Addr)
		http.HandleFunc("/health", middleware.Chain(Health(db, l), middleware.OnlyMethod("GET"), middleware.Logging(l)))
		serverErrors <- apiSrv.ListenAndServe()
	}(l)

	select {
	case err := <-serverErrors:
		return err
	case sig := <-shutdown:
		timeout, _ := time.ParseDuration("5s")
		log.Printf("got: %v : Start shutdown with timeout %s", sig, timeout)
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := apiSrv.Shutdown(ctx)
		if err != nil {
			log.Printf("graceful shutdown timeout in %v : %v", timeout, err)
			err = apiSrv.Close()
		}

		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return err
		}
	}

	return nil
}

func short(s string) string {
	if len(s) > 7 {
		return s[0:7]
	}
	return s
}
