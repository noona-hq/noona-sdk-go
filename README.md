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

Client credentials given in `GetOAuthTokenParams` are sent as HTTP Basic in the `Authorization` header, never in the URL. The noona-api token endpoint documents the header as the preferred method.

```go
c, err := noona.NewAnonymous()
if err != nil {
	panic(errors.Wrap(err, "failed to create client"))
}

resp, err := c.GetOAuthTokenWithResponse(context.Background(), &noona.GetOAuthTokenParams{
	ClientId:     &clientID,
	ClientSecret: &clientSecret,
}, noona.GetOAuthTokenJSONRequestBody{
	GrantType: noona.AuthorizationCode,
	Code:      &code,
})
```
