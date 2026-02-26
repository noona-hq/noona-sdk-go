# Noona SDK for Go

The Noona SDK for Go is an HTTP client library for interacting with the Noona API (`api.noona.is`). It is generated from the Noona OpenAPI spec using `oapi-codegen`.

## Tech Stack

- **Language:** Go 1.24.0
- **Key Dependencies:**
  - `github.com/deepmap/oapi-codegen` — OpenAPI code generation
  - `github.com/pkg/errors` — structured error handling

## Architecture / How it works

- **OAPI-Codegen:** The client is auto-generated from the Noona OpenAPI specification. Do not manually edit generated files.
- **HTTP Client:** Makes typed HTTP calls to `https://api.noona.is`. Not a web framework — contains no server-side logic.
- **Error Handling:** Structured error handling via `github.com/pkg/errors`.

## Key Interfaces / API

- **Client:** Provides typed methods for all Noona API endpoints (e.g. `GetUserWithResponse`, etc.).

## Installation

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

const noonaToken = "<your-noona-token>"

func main() {
	c, err := noona.New(noonaToken)
	if err \!= nil {
		panic(errors.Wrap(err, "failed to create client"))
	}

	resp, err := c.GetUserWithResponse(context.Background(), &noona.GetUserParams{})
	if err \!= nil {
		panic(errors.Wrap(err, "failed to get user"))
	}

	if resp.StatusCode() \!= http.StatusOK {
		panic(errors.Errorf("failed to get user: %d", resp.StatusCode()))
	}

	fmt.Println(*resp.JSON200.Email)
}
```

