package graph

import (
	"database/sql"
	"graphql-service/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	Db *sql.DB
	A  *auth.Auth
}
