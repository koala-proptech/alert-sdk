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
		Type            int32    `json:"type"`
		InstanceID      string   `json:"instance_id"`
		Id              string   `json:"id"`
		ClickAction     string   `json:"click_action"`
	}
)

func (c *Client) MultipleDevice(ctx context.Context, req MultipleDeviceRequest) (*Response, error) {
	url := fmt.Sprintf("%s/push-notif/send", c.url)
	return c.walk(http.MethodPost, url, c.token, req)
}
