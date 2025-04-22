# OpenAPI Specification Guide

This guide explains how to write OpenAPI specifications for the Raybot API.


## Directory Structure

```
api/openapi/
├── components/
│   └── parameters/      # Reusable parameter definitions
│   └── schemas/         # Reusable schema definitions
├── paths/               # API endpoint definitions
└── openapi.yml          # Main OpenAPI document
```


## Writing Specifications

### Main OpenAPI Document

```yaml
openapi: 3.0.0
info:
  version: 0.1.0
  title: Raybot API
  description: Brief description of your API
  license:
    url: https://opensource.org/licenses/MIT
    name: MIT
servers:
  - url: /api/v1
paths:
  /your/path:
    $ref: "./paths/your@path.yml"
```


### Path Files

Create files in the `paths/` directory using `@` instead of `/` in filenames:

```yaml
# paths/resource@id.yml
get:
  summary: Short summary
  operationId: uniqueOperationId
  description: Detailed description
  tags:
    - resourceCategory
  responses:
    "200":
      description: Success response
      content:
        application/json:
          schema:
            $ref: "../components/schemas/your_schema.yml#/ResponseType"
    "400":
      description: Bad Request
      content:
        application/json:
          schema:
            $ref: "../components/schemas/error.yml#/ErrorResponse"
```

### Schema Files

Create reusable schemas in `components/schemas/`:

```yaml
# components/schemas/your_schema.yml
ResponseType:
  type: object
  properties:
    id:
      type: string
      description: Resource identifier
      example: "123e4567-e89b-12d3-a456-426614174000"
      x-order: 1
  required:
    - id
```

## Error Handling

All endpoints should return standard error responses:

```yaml
"400":
  description: Bad Request
  content:
    application/json:
      schema:
        $ref: "../components/schemas/error.yml#/ErrorResponse"
```

The error schema includes code, message, and optional field-specific errors.

## Best Practices

1. Use `x-order` to control property display order
2. Include examples for all properties
3. Use descriptive `operationId` values
4. Reference schemas with `$ref` for reusability
5. Use consistent naming conventions
6. Document all possible response codes
7. Use enum for properties with fixed values
8. Add `x-go-type` when needed for Go code generation
