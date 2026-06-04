package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHandlerJSON verifies the root handler returns {"hello":"world"} JSON.
// (The previous SimpleFactory test was removed when the handler switched to a
// plain map-based payload.)
func TestHandlerJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	var got map[string]string
	if err := json.Unmarshal([]byte(strings.TrimSpace(string(body))), &got); err != nil {
		t.Fatalf("response is not valid JSON: %v (body=%q)", err, string(body))
	}
	if got["hello"] != "world" {
		t.Errorf("hello = %q, want %q", got["hello"], "world")
	}
}
