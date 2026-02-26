# Noona SDK for Go

The Noona SDK for Go provides a set of tools and utilities to interact with Noona services using the Go programming language.

## **Tech Stack**
- **Language:** Go 1.24.0
- **Frameworks:** Gin, Pongo2
- **Key Dependencies:** 
  - github.com/deepmap/oapi-codegen
  - github.com/pkg/errors
  - github.com/gin-gonic/gin
  - github.com/go-playground/validator/v10

## **Architecture / How it works**
- **Modular Design:** The SDK is organized into modules for easy integration and scalability.
- **OAPI-Codegen:** Utilizes OpenAPI code generation for type-safe API interactions.
- **Error Handling:** Implements structured error handling using github.com/pkg/errors.
- **Middleware Support:** Integrates with Gin middleware for enhanced request processing.

## **Key Interfaces / API**
- **Client Interface:** Provides methods to interact with Noona's core services.
- **Event Subscriptions:** Supports subscribing to Noona's event streams for real-time data processing.
- **Data Validation:** Includes validation utilities for ensuring data integrity before API calls.

## **Dependencies**
- **Noona Services:** 
  - Noona Authentication Service
  - Noona Data Processing Service
  - Noona Event Bus