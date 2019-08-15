package sdk

import (
	"context"
	"fmt"
	"net/http"
)

type (
	OTPGenerateRequest struct {
		Destination    string `json:"destination"`
		SenderName     string `json:"sender_name"`
		ProductName    string `json:"product_name"`
		CodeLength     uint8  `json:"code_length,omitempty"`
		CodeValidity   uint16 `json:"code_validity,omitempty"`
		CodeType       string `json:"code_type,omitempty"`
		Template       string `json:"template,omitempty"`
		ResendInterval uint16 `json:"resend_interval,omitempty"`
	}
	OTPValidateRequest struct {
		Uid  string `json:"uid"`
		Code string `json:"code"`
		Ip   string `json:"ip"`
	}
)

func (c *Client) OtpGenerate(ctx context.Context, req OTPGenerateRequest) (*Response, error) {
	url := fmt.Sprintf("%s/otp/generate", c.url)
	return c.walk(http.MethodPost, url, c.token, req)
}

func (c *Client) OtpValidate(ctx context.Context, req OTPValidateRequest) (*Response, error) {
	url := fmt.Sprintf("%s/otp/verify", c.url)
	return c.walk(http.MethodPost, url, c.token, req)
}
