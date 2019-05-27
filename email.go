package sdk

import (
	"context"
	"fmt"
	"net/http"
)

type (
	EmailRequest struct {
		From    string      `json:"from"`
		To      []string    `json:"to"`
		Subject string      `json:"subject"`
		Body    interface{} `json:"body"`
	}
)

func (c *Client) VerificationEmail(ctx context.Context, req EmailRequest) (*Response, error) {
	url := fmt.Sprintf("%s/email/verification", c.url)
	return c.walk(http.MethodPost, url, c.token, req)
}
