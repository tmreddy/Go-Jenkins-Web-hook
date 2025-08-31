package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	if body != "Hello, World!" {
		t.Errorf("expected body 'Hello, World!', got '%s'", body)
	}
}
