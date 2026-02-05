# Golang API Standard HTTP Lib

A simple Go API built with Gin framework following clean architecture patterns.

## Environment Variables

Create a `.env` file in the root directory with the following variables:

- `GIN_MODE` - Gin framework mode (debug/release)
- `FRONTEND_APP_URL` - Frontend application URL for CORS
- `POSTGRES_CONNECT_URL` - PostgreSQL database connection string

## Installation

```bash
go mod download
```

## Running the API

```bash
go run cmd/api/main.go
```

## API Endpoints

The API provides user management endpoints following RESTful patterns.