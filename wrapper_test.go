package hq

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOAuthClientCredentialsSentAsBasicAuth(t *testing.T) {
	id, secret := "id", "secret"
	const basic = "Basic aWQ6c2VjcmV0"
	body := GetOAuthTokenJSONRequestBody{GrantType: RefreshToken, RefreshToken: &secret}

	type captured struct {
		auth, query, body string
	}

	for _, tc := range []struct {
		name     string
		token    string
		params   *GetOAuthTokenParams
		wantAuth string
	}{
		{name: "anonymous with params", params: &GetOAuthTokenParams{ClientId: &id, ClientSecret: &secret}, wantAuth: basic},
		{name: "bearer replaced with params", token: "tok", params: &GetOAuthTokenParams{ClientId: &id, ClientSecret: &secret}, wantAuth: basic},
		{name: "anonymous without params", params: &GetOAuthTokenParams{}, wantAuth: ""},
		{name: "bearer kept without params", token: "tok", params: &GetOAuthTokenParams{}, wantAuth: "Bearer tok"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got captured
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, _ := io.ReadAll(r.Body)
				got = captured{auth: r.Header.Get(NoonaAuthHeader), query: r.URL.RawQuery, body: string(b)}
			}))
			t.Cleanup(server.Close)

			opts := ClientOptions{BaseURL: server.URL}
			c, err := NewAnonymous(opts)
			if tc.token != "" {
				c, err = New(tc.token, opts)
			}
			if err != nil {
				t.Fatal(err)
			}

			if _, err := c.GetOAuthTokenWithResponse(context.Background(), tc.params, body); err != nil {
				t.Fatal(err)
			}

			if got.auth != tc.wantAuth {
				t.Fatalf("Authorization = %q, want %q", got.auth, tc.wantAuth)
			}
			if got.query != "" {
				t.Fatalf("query = %q, want empty", got.query)
			}
			if want := `{"grant_type":"refresh_token","refresh_token":"secret"}`; got.body != want {
				t.Fatalf("body = %q, want %q", got.body, want)
			}
		})
	}

	t.Run("other requests untouched", func(t *testing.T) {
		var auth string
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth = r.Header.Get(NoonaAuthHeader)
		}))
		t.Cleanup(server.Close)

		c, err := New("tok", ClientOptions{BaseURL: server.URL})
		if err != nil {
			t.Fatal(err)
		}
		if _, err := c.GetAdPricingWithResponse(context.Background(), "company"); err != nil {
			t.Fatal(err)
		}
		if auth != "Bearer tok" {
			t.Fatalf("Authorization = %q, want Bearer tok", auth)
		}
	})
}
