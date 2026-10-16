# noona-sdk-go

**Type:** Client library (Go SDK)
**Language:** Go
**Purpose:** Auto-generated Go HTTP client for the Noona API (`api.noona.is`), used by Noona App integrations built in Go.

## Responsibilities
- Provide typed Go methods for all Noona API endpoints
- Handle OAuth token auth for Noona App integrations
  - Client credentials passed in `GetOAuthTokenParams` are sent as HTTP Basic in the `Authorization` header, never in the URL (done by a request editor in `wrapper.go`)
- Generated from Noona OpenAPI spec via oapi-codegen — do not manually edit generated files

## Tech Stack
- **Key deps:** oapi-codegen, github.com/pkg/errors

## Integrations
- **noona-api** — wraps the Noona REST API

## Consumers
app-google-analytics, app-google-reserve, app-storyous, and any other Go-based Noona Apps

## Route here for
- Missing or incorrect Noona API methods in the Go SDK
- SDK regeneration after Noona API spec changes
