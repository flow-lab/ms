package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	database "github.com/flow-lab/ms/internal/platform"
	"log"
	"net/http"
)

type status struct {
	DB bool `json:"db"`
}

// Health checks health status of this application. Will run check against db, if
// everything looks ok then it will return 200 code, 500 otherwise.
func Health(db *sql.DB, log *log.Logger) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		if err := database.StatusCheck(context.Background(), db); err != nil {
			http.Error(w, errMsg(err, log), http.StatusInternalServerError)
			return
		}

		b, _ := json.Marshal(&status{
			DB: true,
		})
		_, _ = fmt.Fprintln(w, string(b))
	}
}

func errMsg(err error, log *log.Logger) string {
	b, _ := json.Marshal(&status{
		DB: false,
	})
	log.Printf("error: %v", err)
	return string(b)
}
