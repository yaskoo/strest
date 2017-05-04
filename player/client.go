package player

import (
	"net/http"

	"github.com/yaskoo/strest/play"
)

// HttpClient is the facace used by Player.
// By default this should just wrap calls to http.Client.
// Useful for testing.
type HttpClient interface {
	Request(ctx *play.Context, s *play.Step, host string) (*http.Request, error)
	Do(req *http.Request) (*http.Response, error)
}

type DefaultHttpClient struct {
	c *http.Client
}

func (dhc *DefaultHttpClient) Request(ctx *play.Context, s *play.Step, host string) (*http.Request, error) {
	str, err := Template(ctx, host+s.Url)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest(s.Method, str, s.Body.Reader())
	if err != nil {
		return nil, err
	}

	for key, value := range s.Headers {
		for _, h := range value.Val {
			str, err = Template(ctx, h)
			if err != nil {
				return nil, err
			}
			req.Header.Add(key, str)
		}
	}
	return req, nil
}

func (dhc *DefaultHttpClient) Do(req *http.Request) (*http.Response, error) {
	return dhc.c.Do(req)
}

func Client() HttpClient {
	return &DefaultHttpClient{
		c: &http.Client{},
	}
}
