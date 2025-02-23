package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"cds-graphql-poc/graph"
	l "cds-graphql-poc/graph/loaders"
	"cds-graphql-poc/graph/model"

	// "cds-graphql-poc/graph/utils"
	"context"
)

// CompaniesByIDOrg is the resolver for the companiesByIdOrg field.
func (r *queryResolver) CompaniesByIDOrg(ctx context.Context, ids []*string) ([]*model.Company, error) {
	// companies, errs := r.CompanyService.GetCompaniesById(ctx, ids)

	// if errs != nil {
	// 	return nil, utils.ConvertErrorsToGqlError(errs)
	// }

	// return companies, nil

	return l.GetCompanies(ctx, ids)
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
