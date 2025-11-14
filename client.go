package landingai

import (
	"context"
	"net/http"
	"time"
)

const (
	// DefaultTimeout is the default timeout for API requests
	DefaultTimeout = 300 * time.Second
)

// Client is the main client for interacting with the Landing AI API
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	region     Region
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// NewClient creates a new Landing AI client with the given API key
func NewClient(apiKey string, opts ...ClientOption) *Client {
	client := &Client{
		apiKey: apiKey,
		region: RegionUS,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	// Set base URL based on region if not already set
	if client.baseURL == "" {
		client.baseURL = client.region.BaseURL()
	}

	return client
}

// WithRegion sets the region for the client
func WithRegion(region Region) ClientOption {
	return func(c *Client) {
		c.region = region
		c.baseURL = region.BaseURL()
	}
}

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets the timeout for API requests
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// Parse initiates a document parsing request
func (c *Client) Parse(ctx context.Context) *ParseRequestBuilder {
	return &ParseRequestBuilder{
		client: c,
		ctx:    ctx,
	}
}

// APIKey returns the API key
func (c *Client) APIKey() string {
	return c.apiKey
}

// BaseURL returns the base URL
func (c *Client) BaseURL() string {
	return c.baseURL
}

// HTTPClient returns the HTTP client
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

// Region returns the region
func (c *Client) Region() Region {
	return c.region
}
