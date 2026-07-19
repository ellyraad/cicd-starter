package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name          string
		authorization string
		wantKey       string
		wantErr       bool
		wantNoHeader  bool
	}

	tests := []test{
		{name: "valid API key", authorization: "ApiKey testing123", wantKey: "testing123"},
		{name: "missing auth header", wantErr: true, wantNoHeader: true},
		{name: "wrong auth key", authorization: "Bearer testing123", wantErr: true},
		{name: "missing API key", authorization: "ApiKey", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := make(http.Header)
			if tt.authorization != "" {
				headers.Set("Authorization", tt.authorization)
			}

			got, err := GetAPIKey(headers)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				if tt.wantNoHeader && !errors.Is(err, ErrNoAuthHeaderIncluded) {
					t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.wantKey {
				t.Errorf("expected API key %q, got %q", tt.wantKey, got)
			}
		})
	}
}
