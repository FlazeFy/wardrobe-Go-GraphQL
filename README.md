# wardrobe-Go-GraphQL

## Prerequisites

- Go 1.25+
- PostgreSQL

## Installation

```bash
git clone <repository>
cd wardrobe-graphql

go mod tidy
gqlgen generate
go run main.go
```

The GraphQL Playground will be available at:

```
http://localhost:8080/playground
```