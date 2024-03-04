package hq

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

const (
	NoonaBaseURL    = "https://api.noona.is"
	NoonaAuthHeader = "HQ-ACCESS-KEY"
)

type ClientOptions struct {
	// Defaults to https://api.noona.is
	BaseURL string
}

func New(token string, options ...ClientOptions) (*ClientWithResponses, error) {
	opts := ClientOptions{
		BaseURL: NoonaBaseURL,
	}

	if len(options) > 0 {
		opts = options[0]
	}

	authHeader := WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add(NoonaAuthHeader, token)
		return nil
	})

	c, err := NewClientWithResponses(opts.BaseURL, authHeader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client")
	}

	return c, nil
}
