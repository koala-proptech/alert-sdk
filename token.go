package sdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) Token(ctx context.Context) (r *Response, err error) {
	url := fmt.Sprintf("%s/token", baseUrl)
	r, err = c.talk(http.MethodGet, url, c.uid, nil)
	if err != nil {
		return
	}
	if nil == r.Content {
		return
	}
	token, ok := r.Content["token"]
	if !ok {
		err = errors.New("no valid token assigned from upstream")
		return
	}
	c.token = token.(string)
	return
}
