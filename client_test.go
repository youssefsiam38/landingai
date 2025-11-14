package landingai

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name   string
		apiKey string
		opts   []ClientOption
	}{
		{
			name:   "basic client",
			apiKey: "test-api-key",
			opts:   nil,
		},
		{
			name:   "client with EU region",
			apiKey: "test-api-key",
			opts:   []ClientOption{WithRegion(RegionEU)},
		},
		{
			name:   "client with custom timeout",
			apiKey: "test-api-key",
			opts:   []ClientOption{WithTimeout(5 * time.Minute)},
		},
		{
			name:   "client with custom base URL",
			apiKey: "test-api-key",
			opts:   []ClientOption{WithBaseURL("https://custom.example.com")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.apiKey, tt.opts...)
			if client == nil {
				t.Error("NewClient() returned nil")
				return
			}
			if client.APIKey() != tt.apiKey {
				t.Errorf("APIKey() = %v, want %v", client.APIKey(), tt.apiKey)
			}
			if client.BaseURL() == "" {
				t.Error("BaseURL() returned empty string")
			}
		})
	}
}

func TestRegion_BaseURL(t *testing.T) {
	tests := []struct {
		name   string
		region Region
		want   string
	}{
		{
			name:   "US region",
			region: RegionUS,
			want:   "https://api.va.landing.ai",
		},
		{
			name:   "EU region",
			region: RegionEU,
			want:   "https://api.va.eu-west-1.landing.ai",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.region.BaseURL(); got != tt.want {
				t.Errorf("Region.BaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIError_Methods(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		message    string
		checkFunc  func(*APIError) bool
		want       bool
	}{
		{
			name:       "IsUnauthorized",
			statusCode: 401,
			message:    "Unauthorized",
			checkFunc:  (*APIError).IsUnauthorized,
			want:       true,
		},
		{
			name:       "IsPaymentRequired",
			statusCode: 402,
			message:    "Payment Required",
			checkFunc:  (*APIError).IsPaymentRequired,
			want:       true,
		},
		{
			name:       "IsRateLimited",
			statusCode: 429,
			message:    "Too Many Requests",
			checkFunc:  (*APIError).IsRateLimited,
			want:       true,
		},
		{
			name:       "IsBadRequest",
			statusCode: 400,
			message:    "Bad Request",
			checkFunc:  (*APIError).IsBadRequest,
			want:       true,
		},
		{
			name:       "IsValidationError",
			statusCode: 422,
			message:    "Unprocessable Entity",
			checkFunc:  (*APIError).IsValidationError,
			want:       true,
		},
		{
			name:       "IsServerError",
			statusCode: 500,
			message:    "Internal Server Error",
			checkFunc:  (*APIError).IsServerError,
			want:       true,
		},
		{
			name:       "IsTimeout",
			statusCode: 504,
			message:    "Gateway Timeout",
			checkFunc:  (*APIError).IsTimeout,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &APIError{
				StatusCode: tt.statusCode,
				Message:    tt.message,
			}
			if got := tt.checkFunc(err); got != tt.want {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
			if err.Error() == "" {
				t.Error("Error() returned empty string")
			}
		})
	}
}
