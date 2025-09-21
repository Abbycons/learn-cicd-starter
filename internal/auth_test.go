package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headerVal string
		wantKey   string
		wantErr   bool
	}{
		{"valid key", "Bearer abc123", "abc123", false},
		{"missing header", "", "", true},
		{"wrong scheme", "Token xyz", "", true},
		{"empty token", "Bearer ", "", true},
		{"extra spaces", "Bearer   spacedToken  ", "spacedToken", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			if tc.headerVal != "" {
				req.Header.Set("Authorization", tc.headerVal)
			}

			got, err := GetAPIKey(req)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error, got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error, got %v", err)
				}
				if got != tc.wantKey {
					t.Errorf("expected key %q, got %q", tc.wantKey, got)
				}
			}
		})
	}
}
