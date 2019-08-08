package sdk

import (
	"context"
	"fmt"
	"net/http"
)

type (
	MultipleDeviceRequest struct {
		RegistrationIDS []string `json:"registration_ids"`
		Title           string   `json:"title"`
		Message         string   `json:"message"`
	}
)

func (c *Client) MultipleDevice(ctx context.Context, req MultipleDeviceRequest) (*Response, error) {
	url := fmt.Sprintf("%s/push-notif/send", c.url)
	return c.walk(http.MethodPost, url, c.token, req)
}
