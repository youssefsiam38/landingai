package landingai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// ParseRequestBuilder is a builder for Parse API requests
type ParseRequestBuilder struct {
	client      *Client
	ctx         context.Context
	model       *string
	documentURL *string
	filePath    string
	fileData    []byte
	fileName    string
	split       *SplitType
}

// WithModel sets the model version to use for parsing
// Examples: "dpt-2-latest", "dpt-2-20250919", "dpt-1-latest", "DPT-2-mini-latest"
func (b *ParseRequestBuilder) WithModel(model string) *ParseRequestBuilder {
	b.model = &model
	return b
}

// WithURL sets the document URL to parse
func (b *ParseRequestBuilder) WithURL(url string) *ParseRequestBuilder {
	b.documentURL = &url
	return b
}

// WithFile sets the file path to upload and parse
func (b *ParseRequestBuilder) WithFile(filePath string) *ParseRequestBuilder {
	b.filePath = filePath
	return b
}

// WithFileData sets the file data directly (with filename)
func (b *ParseRequestBuilder) WithFileData(data []byte, filename string) *ParseRequestBuilder {
	b.fileData = data
	b.fileName = filename
	return b
}

// WithSplit enables document splitting at the specified level
func (b *ParseRequestBuilder) WithSplit(split SplitType) *ParseRequestBuilder {
	b.split = &split
	return b
}

// WithPageSplit is a convenience method to enable page-level splitting
func (b *ParseRequestBuilder) WithPageSplit() *ParseRequestBuilder {
	split := SplitTypePage
	b.split = &split
	return b
}

// Do executes the parse request
func (b *ParseRequestBuilder) Do() (*ParseResponse, error) {
	// Validate inputs
	if b.documentURL != nil && (b.filePath != "" || b.fileData != nil) {
		return nil, fmt.Errorf("cannot provide both document URL and file")
	}
	if b.documentURL == nil && b.filePath == "" && b.fileData == nil {
		return nil, fmt.Errorf("must provide either document URL or file")
	}

	// Create the request
	req, err := b.buildRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	// Execute the request
	resp, err := b.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Handle errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, b.handleErrorResponse(resp.StatusCode, body)
	}

	// Parse successful response
	var parseResp ParseResponse
	if err := json.Unmarshal(body, &parseResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &parseResp, nil
}

// buildRequest constructs the HTTP request
func (b *ParseRequestBuilder) buildRequest() (*http.Request, error) {
	url := fmt.Sprintf("%s/v1/ade/parse", b.client.baseURL)

	var req *http.Request
	var err error

	if b.documentURL != nil {
		// URL-based request
		req, err = b.buildURLRequest(url)
	} else {
		// File-based request
		req, err = b.buildFileRequest(url)
	}

	if err != nil {
		return nil, err
	}

	// Add authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.client.apiKey))

	// Set context
	req = req.WithContext(b.ctx)

	return req, nil
}

// buildURLRequest builds a request with document_url
func (b *ParseRequestBuilder) buildURLRequest(url string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add document_url
	if err := writer.WriteField("document_url", *b.documentURL); err != nil {
		return nil, err
	}

	// Add optional fields
	if b.model != nil {
		if err := writer.WriteField("model", *b.model); err != nil {
			return nil, err
		}
	}
	if b.split != nil {
		if err := writer.WriteField("split", string(*b.split)); err != nil {
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(b.ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

// buildFileRequest builds a request with file upload
func (b *ParseRequestBuilder) buildFileRequest(url string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Determine file data and name
	var fileData []byte
	var fileName string
	var err error

	if b.fileData != nil {
		fileData = b.fileData
		fileName = b.fileName
	} else {
		// Read file from path
		fileData, err = os.ReadFile(b.filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		fileName = filepath.Base(b.filePath)
	}

	// Add file
	part, err := writer.CreateFormFile("document", fileName)
	if err != nil {
		return nil, err
	}
	if _, err = part.Write(fileData); err != nil {
		return nil, err
	}

	// Add optional fields
	if b.model != nil {
		err = writer.WriteField("model", *b.model)
		if err != nil {
			return nil, err
		}
	}
	if b.split != nil {
		err = writer.WriteField("split", string(*b.split))
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(b.ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

// handleErrorResponse processes error responses from the API
func (b *ParseRequestBuilder) handleErrorResponse(statusCode int, body []byte) error {
	// Try to parse as validation error
	if statusCode == StatusUnprocessableEntity {
		var valErr ValidationErrors
		if err := json.Unmarshal(body, &valErr); err == nil {
			return &valErr
		}
	}

	// Create generic API error
	apiErr := &APIError{
		StatusCode: statusCode,
		Message:    getErrorMessage(statusCode),
	}

	// Try to extract detail from body
	var errorDetail map[string]interface{}
	if err := json.Unmarshal(body, &errorDetail); err == nil {
		if detail, ok := errorDetail["detail"]; ok {
			apiErr.Detail = detail
		}
	} else {
		// If JSON parsing fails, use raw body as detail
		apiErr.Detail = string(body)
	}

	return apiErr
}

// getErrorMessage returns a user-friendly error message for common status codes
func getErrorMessage(statusCode int) string {
	switch statusCode {
	case StatusBadRequest:
		return "Bad request: Invalid request parameters"
	case StatusUnauthorized:
		return "Unauthorized: Invalid or missing API key"
	case StatusPaymentRequired:
		return "Payment required: Insufficient credits"
	case StatusUnprocessableEntity:
		return "Unprocessable entity: Input validation failed"
	case StatusTooManyRequests:
		return "Too many requests: Rate limit exceeded"
	case StatusInternalServerError:
		return "Internal server error: Failed to process document"
	case StatusGatewayTimeout:
		return "Gateway timeout: Request processing exceeded time limit"
	default:
		return fmt.Sprintf("API request failed with status %d", statusCode)
	}
}
