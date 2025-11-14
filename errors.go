package landingai

import "fmt"

// APIError represents an error returned by the Landing AI API
type APIError struct {
	StatusCode int
	Message    string
	Detail     interface{}
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Detail != nil {
		return fmt.Sprintf("Landing AI API error (status %d): %s - %v", e.StatusCode, e.Message, e.Detail)
	}
	return fmt.Sprintf("Landing AI API error (status %d): %s", e.StatusCode, e.Message)
}

// IsUnauthorized returns true if the error is due to invalid authentication
func (e *APIError) IsUnauthorized() bool {
	return e.StatusCode == StatusUnauthorized
}

// IsPaymentRequired returns true if the error is due to insufficient credits
func (e *APIError) IsPaymentRequired() bool {
	return e.StatusCode == StatusPaymentRequired
}

// IsRateLimited returns true if the error is due to rate limiting
func (e *APIError) IsRateLimited() bool {
	return e.StatusCode == StatusTooManyRequests
}

// IsBadRequest returns true if the error is due to invalid request parameters
func (e *APIError) IsBadRequest() bool {
	return e.StatusCode == StatusBadRequest
}

// IsValidationError returns true if the error is due to input validation failure
func (e *APIError) IsValidationError() bool {
	return e.StatusCode == StatusUnprocessableEntity
}

// IsServerError returns true if the error is due to a server-side failure
func (e *APIError) IsServerError() bool {
	return e.StatusCode >= StatusInternalServerError
}

// IsTimeout returns true if the error is due to a timeout
func (e *APIError) IsTimeout() bool {
	return e.StatusCode == StatusGatewayTimeout
}

// IsPartialContent returns true if some pages failed but parsing succeeded
func (e *APIError) IsPartialContent() bool {
	return e.StatusCode == StatusPartialContent
}

// ValidationError represents a validation error from the API
type ValidationError struct {
	Location []interface{} `json:"loc"`
	Message  string        `json:"msg"`
	Type     string        `json:"type"`
}

// ValidationErrors represents a collection of validation errors
type ValidationErrors struct {
	Detail []ValidationError `json:"detail"`
}

// Error implements the error interface
func (v *ValidationErrors) Error() string {
	if len(v.Detail) == 0 {
		return "validation error"
	}
	return fmt.Sprintf("validation error: %s", v.Detail[0].Message)
}
