# The Noona Golang SDK

Noona SDK for the Go programming language.

## Installing

```bash
go get github.com/noona-hq/noona-sdk-go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	noona "github.com/noona-hq/noona-sdk-go"
	"github.com/pkg/errors"
)

const (
	noonaToken       = "<your-noona-token>"
)

func main() {
	c, err := noona.New(noonaToken)
	if err != nil {
		panic(errors.Wrap(err, "failed to create client"))
	}

	resp, err := c.GetUserWithResponse(context.Background(), &noona.GetUserParams{})
	if err != nil {
		panic(errors.Wrap(err, "failed to get user"))
	}

	if resp.StatusCode() != http.StatusOK {
		panic(errors.Errorf("failed to get user: %d", resp.StatusCode()))
	}

	user := resp.JSON200

	// Do something with user
	fmt.Println(*user.Email)
}


```

## OAuth token exchange

Noona Apps exchange an authorization code or refresh token for an access token using the app's client credentials. `NewWithClientCredentials` sends them as HTTP Basic on every request. Client credentials are only accepted in the `Authorization` header, never in the URL or request body.

```go
c, err := noona.NewWithClientCredentials(noona.OAuthClientCredentials{
	ClientID:     "<your-client-id>",
	ClientSecret: "<your-client-secret>",
})
if err != nil {
	panic(errors.Wrap(err, "failed to create client"))
}

code := "<authorization-code>"
resp, err := c.GetOAuthTokenWithResponse(context.Background(), noona.GetOAuthTokenJSONRequestBody{
	GrantType: noona.AuthorizationCode,
	Code:      &code,
})
if err != nil {
	panic(errors.Wrap(err, "failed to exchange code"))
}

token := resp.JSON200

refreshToken := *token.RefreshToken
resp, err = c.GetOAuthTokenWithResponse(context.Background(), noona.GetOAuthTokenJSONRequestBody{
	GrantType:    noona.RefreshToken,
	RefreshToken: &refreshToken,
})
```
