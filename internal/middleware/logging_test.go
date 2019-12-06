package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLogging(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	l := log.New(os.Stdout, "", 0)
	h := Chain(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}, Logging(l))

	h.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("response code was %v instead of 200", rec.Code)
	}
}
