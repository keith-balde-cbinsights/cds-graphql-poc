package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"cds-graphql-poc/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CompanyService service.CompanyService
}
