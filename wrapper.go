package hq

import (
	"context"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const (
	NoonaBaseURL              = "https://api.noona.is"
	NoonaAuthHeader           = "Authorization"
	NoonaAcceptLanguageHeader = "Accept-Language"
	NoonaDefaultLanguage      = "en"

	oauthTokenPathSuffix   = "/v1/hq/oauth/token"
	oauthClientIDParam     = "client_id"
	oauthClientSecretParam = "client_secret"
)

type ClientOptions struct {
	// Defaults to https://api.noona.is
	BaseURL        string
	AcceptLanguage string
}

func moveOAuthClientCredentialsToHeader(ctx context.Context, req *http.Request) error {
	if req.Method != http.MethodPost || !strings.HasSuffix(req.URL.Path, oauthTokenPathSuffix) {
		return nil
	}

	q := req.URL.Query()
	if !q.Has(oauthClientIDParam) && !q.Has(oauthClientSecretParam) {
		return nil
	}

	id, secret := q.Get(oauthClientIDParam), q.Get(oauthClientSecretParam)
	q.Del(oauthClientIDParam)
	q.Del(oauthClientSecretParam)
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(id, secret)

	return nil
}

func New(token string, options ...ClientOptions) (*ClientWithResponses, error) {
	opts := ClientOptions{
		BaseURL:        NoonaBaseURL,
		AcceptLanguage: NoonaDefaultLanguage,
	}

	if len(options) > 0 {
		opts = options[0]
	}

	authHeader := WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add(NoonaAuthHeader, "Bearer "+token)
		req.Header.Add(NoonaAcceptLanguageHeader, opts.AcceptLanguage)
		return nil
	})

	c, err := NewClientWithResponses(opts.BaseURL, authHeader, WithRequestEditorFn(moveOAuthClientCredentialsToHeader))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client")
	}

	return c, nil
}

func NewAnonymous(options ...ClientOptions) (*ClientWithResponses, error) {
	opts := ClientOptions{
		BaseURL: NoonaBaseURL,
	}

	if len(options) > 0 {
		opts = options[0]
	}

	c, err := NewClientWithResponses(opts.BaseURL, WithRequestEditorFn(moveOAuthClientCredentialsToHeader))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client")
	}

	return c, nil
}
