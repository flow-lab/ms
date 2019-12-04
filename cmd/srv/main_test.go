package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func ok(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func TestLogging(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	check(t, err)

	rec := httptest.NewRecorder()
	h := Chain(ok, Logging())

	h.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("response code was %v instead of 200", rec.Code)
	}
}

func TestOnlyMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	check(t, err)

	rec := httptest.NewRecorder()
	h := Chain(ok, OnlyMethod("GET"))

	h.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("response code was %v instead of 200", rec.Code)
	}
}

func TestOnlyMethod2(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	check(t, err)

	rec := httptest.NewRecorder()
	h := Chain(ok, OnlyMethod("GET"))

	h.ServeHTTP(rec, req)

	if rec.Code != 400 {
		t.Errorf("expecting 400. Got %v", rec.Code)
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
