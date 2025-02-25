package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"cds-graphql-poc/middleware/contextcache"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// CompanyService company.Service
	Cache *contextcache.Cache
}
