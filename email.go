package sdk

import (
	"context"
	"fmt"
	"net/http"
)

type Attachment struct {
	AttachmentURL      string `json:"attachment_url"`
	AttachmentFileName string `json:"attachment_filename"`
}

type (
	EmailRequest struct {
		From               string       `json:"from"`
		To                 []string     `json:"to"`
		Subject            string       `json:"subject"`
		Body               interface{}  `json:"body"`
		AttachmentURL      string       `json:"attachment_url"`
		AttachmentFileName string       `json:"attachment_filename"`
		Attachments        []Attachment `json:"attachments"`
	}
)

func (c *Client) VerificationEmail(ctx context.Context, req EmailRequest) (*Response, error) {
	url := fmt.Sprintf("%s/email/verification", c.url)
	return c.walk(ctx, http.MethodPost, url, c.token, req)
}

func (c *Client) SendEmail(ctx context.Context, req EmailRequest) (*Response, error) {
	url := fmt.Sprintf("%s/email", c.url)
	return c.walk(ctx, http.MethodPost, url, c.token, req)
}

func (c *Client) WithAttachment(ctx context.Context, req EmailRequest) (*Response, error) {
	url := fmt.Sprintf("%s/email/with-attachment", c.url)
	return c.walk(ctx, http.MethodPost, url, c.token, req)
}

func (c *Client) LoanSubmissionBank(ctx context.Context, req EmailRequest) (*Response, error) {
	url := fmt.Sprintf("%s/email/loan/bank", c.url)
	return c.walk(ctx, http.MethodPost, url, c.token, req)
}

func (c *Client) LoanSubmissionCustomer(ctx context.Context, req EmailRequest) (*Response, error) {
	url := fmt.Sprintf("%s/email/loan/customer", c.url)
	return c.walk(ctx, http.MethodPost, url, c.token, req)
}

func (c *Client) WelcomeEmail(ctx context.Context, req EmailRequest) (*Response, error) {
	url := fmt.Sprintf("%s/email/welcome", c.url)
	return c.walk(ctx, http.MethodPost, url, c.token, req)
}
