package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

var (
	userAgent           = "alert-sdk:v0.0.1"
	mimeApplicationJson = "application/json"
	baseUrl             = "http://localhost:8080"
)

type (
	httpClient interface {
		Do(*http.Request) (*http.Response, error)
	}
	Option func(*Client)
	Client struct {
		uid, token string
		httpClient httpClient
	}
	ErrorResponse struct {
		Code    string            `json:"code"`
		Message string            `json:"message"`
		Reasons map[string]string `json:"reasons"`
	}
	Response struct {
		RequestID string                 `json:"request_id"`
		Status    int                    `json:"status"`
		Content   map[string]interface{} `json:"content,omitempty"`
		Error     *ErrorResponse         `json:"error,omitempty"`
	}
)

func SetBaseUrl(url string) {
	baseUrl = url
}

func OptionHTTPClient(client httpClient) func(*Client) {
	return func(c *Client) {
		c.httpClient = client
	}
}

func New(uid string, options ...Option) *Client {
	s := &Client{
		uid:        uid,
		httpClient: &http.Client{},
	}
	s.AddOptions(options...)
	return s
}

func (c *Client) AddOptions(options ...Option) {
	for _, opt := range options {
		opt(c)
	}
}

func (c *Client) build(method, url, token string, payload io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, errors.WithMessage(err, "failed creating request")
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	r.Header.Set("Content-Type", mimeApplicationJson)
	r.Header.Set("Accept", mimeApplicationJson)
	r.Header.Set("Cache-Control", "no-cache")
	r.Header.Set("User-Agent", userAgent)
	return r, nil
}

func (c *Client) request(r *http.Request) (*Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, errors.WithMessage(err, "failed communicating with upstream")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
		return &Response{Status: resp.StatusCode}, errors.New(http.StatusText(resp.StatusCode))
	}

	var rs Response
	if err := json.NewDecoder(resp.Body).Decode(&rs); err != nil {
		return nil, errors.WithMessage(err, "failed decoding response")
	}

	return &rs, nil
}

func (c *Client) walk(method, url, token string, payload interface{}) (resp *Response, err error) {
	// first request - exchange token before proceed
	if len(c.token) == 0 {
		_, err = c.Token(context.Background())
		if err != nil {
			return
		}
		return c.talk(method, url, c.token, payload)
	}

	// token intact, lets attempt to process
	resp, err = c.talk(method, url, token, payload)
	if err != nil {
		return
	}

	// ouch! probably expired, we need to refresh
	if http.StatusUnauthorized == resp.Status {
		c.token = ""
		return c.walk(method, url, token, payload)
	}
	return
}

func (c *Client) talk(method, url, token string, payload interface{}) (*Response, error) {
	var ir io.Reader
	if nil != payload {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.WithMessage(err, "failed encoding request payload")
		}
		ir = bytes.NewReader(b)
	}
	r, err := c.build(method, url, token, ir)
	if err != nil {
		return nil, err
	}
	return c.request(r)
}
