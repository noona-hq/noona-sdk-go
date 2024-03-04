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
	noonaAPIBasePath = "https://api.noona.is"
	noonaAuthHeader  = "HQ-ACCESS-KEY"
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
