package growwapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client to access groww apis
type Client struct {
	accessToken string
	httpClient  *http.Client
}

// NewClient creates a new Client
func NewClient(accessToken string, httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return Client{accessToken, httpClient}
}

// ErrorCode are codes returned by GROWW APIs
type ErrorCode string

const (
	ErrorCodeGA000 ErrorCode = "GA000"
	ErrorCodeGA001 ErrorCode = "GA001"
	ErrorCodeGA003 ErrorCode = "GA003"
	ErrorCodeGA004 ErrorCode = "GA004"
	ErrorCodeGA005 ErrorCode = "GA005"
	ErrorCodeGA006 ErrorCode = "GA006"
	ErrorCodeGA007 ErrorCode = "GA007"
)

var errorCodeMessages = map[ErrorCode]string{
	ErrorCodeGA000: "Internal error occurred",
	ErrorCodeGA001: "Bad request",
	ErrorCodeGA003: "Unable to serve request currently",
	ErrorCodeGA004: "Requested entity does not exist",
	ErrorCodeGA005: "User not authorised to perform this operation",
	ErrorCodeGA006: "Requested entity does not exist",
	ErrorCodeGA007: "Duplicate order reference id",
}

// Message returns the message related to ErrorCode.
// If you want the actual message returned by the API, use Error.Message
func (e ErrorCode) Message() string {
	if got, ok := errorCodeMessages[e]; ok {
		return got
	}

	return "Unknown Error"
}

type errorResponse struct {
	Status string `json:"status"`
	Error  Error  `json:"error"`
}

// Error represents the error data returned from Groww APIs
type Error struct {
	Code     ErrorCode `json:"code"`
	Message  string    `json:"message"`
	Metadata any       `json:"metadata"`
}

type apiResponse[T any] struct {
	Status  string `json:"status"`
	Payload T      `json:"payload"`
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (c *Client) headers() http.Header {
	headers := make(http.Header)

	headers.Add("Accept", "application/json")
	headers.Add("Content-Type", "application/json")
	headers.Add("X-API-VERSION", "1.0")
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	return headers
}

type asQueryParam interface {
	queryParams() url.Values
}

func doGetRequest[T any](c *Client, url string, queries asQueryParam) (out T, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return out, fmt.Errorf("http.NewRequest: %w", err)
	}

	params := queries.queryParams()
	req.URL.RawQuery = params.Encode()

	return doRequest[T](c, req)
}

func doPostRequest[T any](c *Client, url string, body any) (out T, err error) {
	msg, err := json.Marshal(body)
	if err != nil {
		return out, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(msg))
	if err != nil {
		return out, fmt.Errorf("http.NewRequest: %w", err)
	}

	return doRequest[T](c, req)
}

func doRequest[T any](c *Client, req *http.Request) (out T, err error) {
	req.Header = c.headers()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return out, fmt.Errorf("c.httpClient.Do: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var r apiResponse[T]
		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			return out, fmt.Errorf("json.NewDecoder(success_response): %w", err)
		}
		return r.Payload, nil

	default:
		var e errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return out, fmt.Errorf("json.NewDecoder(success_response): %w", err)
		}
		return out, e.Error
	}
}
