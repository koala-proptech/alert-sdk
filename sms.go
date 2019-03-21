package sdk

import (
	"context"
	"fmt"
	"net/http"
)

type (
	SingleSmsRequest struct {
		Destination string `json:"destination"`
		SenderName  string `json:"sender_name"`
		Text        string `json:"text"`
	}
)

func (c *Client) SingleSMS(ctx context.Context, req SingleSmsRequest) (*Response, error) {
	url := fmt.Sprintf("%s/sms/single", c.url)
	return c.walk(http.MethodPost, url, c.token, req)
}
