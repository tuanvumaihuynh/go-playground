# OpenAPI Code Generation with Go

This project demonstrates how to use OpenAPI specifications to generate Go server code automatically using the `oapi-codegen` tool. It implements a RESTful API for user management with standard CRUD operations.

## Project Structure

```
oapicodegen/
├── cmd/
│   └── main.go               # Application entry point
├── internal/
│   └── http/                 # HTTP server implementation
│       ├── gen/              # Generated API code
│       ├── service.go        # HTTP service setup
│       ├── swagger/          # Swagger UI integration
│       └── user_handler.go   # User endpoint handlers
├── openapi/                  # OpenAPI specifications
│   ├── components/           # Reusable schema definitions
│   ├── paths/                # API endpoint definitions
│   ├── embedded.go           # Embedded spec for Swagger UI
│   └── openapi.yml           # Main OpenAPI document
└── pkg/
    └── cmdutil/              # Command utilities (signal handling)
```

## Key Features

- **OpenAPI-first API design** - Define your API in YAML before writing any code
- **Automatic code generation** - Generate server stubs, models, and validators from OpenAPI specs
- **Type-safe API implementation** - Compile-time checking of API implementation correctness
- **Standard library HTTP server** - Uses Go's built-in `net/http` package without external frameworks
- **Swagger UI integration** - Interactive API documentation and testing
- **Graceful shutdown** - Proper handling of OS signals for clean application termination

## Generated Components

The `oapi-codegen` tool generates several important components:

1. **Data models** - Go structs that match your OpenAPI schemas
2. **Server interfaces** - Interfaces that your handlers must implement
3. **Request/response types** - Type-safe request and response objects
4. **Validation** - Path/query parameter validation and binding

## How It Works

1. API is defined in OpenAPI specification files:
   - `openapi.yml` - The main OpenAPI document
   - `paths/*.yml` - Individual endpoint definitions
   - `components/schemas/*.yml` - Reusable schemas (data models)

2. The bundled specification is generated with Redocly CLI:
   ```bash
   redocly bundle openapi/openapi.yml --output openapi/openapi.gen.yml --ext yml
   ```

3. Go code is generated from the bundled specification:
   ```bash
   go tool oapi-codegen -config oapi-codegen.yml openapi/openapi.gen.yml
   ```

4. API handlers are implemented by satisfying the generated interface in `user_handler.go`

5. The HTTP service is set up in `service.go` to use the generated code

## API Endpoints

| Method | Path           | Description            | Response |
|--------|----------------|------------------------|----------|
| GET    | /users         | List all users         | 200: UserListResponse<br>400: ErrorResponse |
| POST   | /users         | Create a new user      | 201: UserResponse<br>400: ErrorResponse |
| GET    | /users/{userId}| Get user by ID         | 200: UserResponse<br>404: ErrorResponse |

## Getting Started

### Prerequisites

- Go 1.24 or higher
- pnpm (for running Redocly CLI)

### Running the Application

1. Generate the OpenAPI code:
   ```bash
   make gen-oapi
   ```

2. Run the server:
   ```bash
   make run
   ```

3. Access the Swagger UI:
   ```
   http://localhost:3000/docs
   ```

## Next Steps & Learning Resources

- Add more API endpoints (DELETE, PUT, PATCH)
- Implement middleware for authentication, logging, etc.
- Connect to a real database instead of using static responses
- Add tests for API handlers

### Resources

- [oapi-codegen Documentation](https://github.com/oapi-codegen/oapi-codegen)
- [OpenAPI Specification](https://swagger.io/specification/)
- [Redocly CLI Documentation](https://redocly.com/docs/cli/)
