package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	resolvers "cds-graphql-poc/graph"
	"cds-graphql-poc/graph/model"
	"context"
	"fmt"
)

// CompaniesByIDOrg is the resolver for the companiesByIdOrg field.
func (r *queryResolver) CompaniesByIDOrg(ctx context.Context, id string) ([]*model.Company, error) {
	panic(fmt.Errorf("not implemented: CompaniesByIDOrg - companiesByIdOrg"))
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
