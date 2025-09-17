package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer testkey123")

	got, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("GetAPIKey() returned unexpected error: %v", err)
	}

	want := "testkey123"
	if got != want {
		t.Errorf("GetAPIKey() = %q; want %q", got, want)
	}
}

