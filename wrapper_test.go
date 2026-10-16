package hq

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewWithClientCredentialsSendsBasicAuth(t *testing.T) {
	var gotAuth, gotLanguage string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get(NoonaAuthHeader)
		gotLanguage = r.Header.Get(NoonaAcceptLanguageHeader)
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	c, err := NewWithClientCredentials(
		OAuthClientCredentials{ClientID: "id", ClientSecret: "secret"},
		ClientOptions{BaseURL: server.URL, AcceptLanguage: "is"},
	)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.GetAdPricingWithResponse(context.Background(), "company"); err != nil {
		t.Fatal(err)
	}

	if want := "Basic aWQ6c2VjcmV0"; gotAuth != want {
		t.Fatalf("Authorization = %q, want %q", gotAuth, want)
	}
	if gotLanguage != "is" {
		t.Fatalf("Accept-Language = %q, want is", gotLanguage)
	}
}
