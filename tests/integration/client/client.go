package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const DefaultEndpoint = "http://127.0.0.1:8080/query"
const EndpointEnvVar = "GRAPHQL_ENDPOINT"

// Request body
type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQL error entry
type GQLError struct {
	Message string `json:"message"`
}

// Raw JSON response from a GraphQL endpoint
type Response struct {
	Data   json.RawMessage `json:"data"`
	Errors []GQLError      `json:"errors"`
}

// Status of response contains GraphQL errors
func (r Response) HasErrors() bool {
	return len(r.Errors) > 0
}

// Get first error message, or "" if none
func (r Response) FirstErrorMessage() string {
	if len(r.Errors) == 0 {
		return ""
	}
	return r.Errors[0].Message
}

// Decodes the response data
func (r Response) Unmarshal(v interface{}) error {
	return json.Unmarshal(r.Data, v)
}

type Client struct {
	endpoint   string
	httpClient *http.Client
	headers    map[string]string
}
type Option func(*Client)

// Overrides the target GraphQL endpoint
func WithEndpoint(url string) Option {
	return func(c *Client) { c.endpoint = url }
}

// Sets a header sent with every request
func WithHeader(key, value string) Option {
	return func(c *Client) {
		if c.headers == nil {
			c.headers = map[string]string{}
		}
		c.headers[key] = value
	}
}

// New builds a Client using DefaultEndpoint, env var, or the given options
func New(opts ...Option) *Client {
	c := &Client{
		endpoint:   resolveEndpoint(),
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Picks the endpoint from env var or falls back to default
func resolveEndpoint() string {
	if v := os.Getenv(EndpointEnvVar); v != "" {
		return v
	}

	return DefaultEndpoint
}

// Sends a GraphQL request and returns the parsed response
func (c *Client) Do(req Request) (Response, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return Response{}, fmt.Errorf("client: marshal request: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.endpoint, bytes.NewBuffer(body))
	if err != nil {
		return Response{}, fmt.Errorf("client: build request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	for k, v := range c.headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return Response{}, fmt.Errorf("client: send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("client: unexpected status %d from %s", resp.StatusCode, c.endpoint)
	}

	var gqlResp Response
	if err := json.NewDecoder(resp.Body).Decode(&gqlResp); err != nil {
		return Response{}, fmt.Errorf("client: decode response: %w", err)
	}

	return gqlResp, nil
}
