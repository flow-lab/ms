package main

import (
	"context"
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
func Health(c database.Config) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		s, err := database.Open(c)
		if err != nil {
			b, _ := json.Marshal(&status{
				DB: false,
			})
			log.Printf("got an error: %v", err)
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		err = database.StatusCheck(context.Background(), s)
		if err != nil {
			log.Printf("got an error: %v", err)
			b, _ := json.Marshal(&status{
				DB: false,
			})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		b, _ := json.Marshal(&status{
			DB: true,
		})
		_, _ = fmt.Fprintln(w, string(b))
	}
}
