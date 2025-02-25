package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"cds-graphql-poc/graph"
	l "cds-graphql-poc/graph/loaders"
	"cds-graphql-poc/graph/model"
	"cds-graphql-poc/utils"

	"context"
)

// CompaniesByIDOrg is the resolver for the companiesByIdOrg field.
func (r *queryResolver) CompaniesByIDOrg(ctx context.Context, ids []*string) ([]*model.Company, error) {
	// companies, errs := r.CompanyService.GetCompaniesById(ctx, ids)

	// if errs != nil {
	// 	return nil, utils.ConvertErrorsToGqlError(errs)
	// }

	// return companies, nil

	intIds, err := utils.ConvertStringsToInts(ids)
	if err != nil {
		return nil, err
	}

	companies, notFound := r.Cache.GetCompanies(ctx, intIds)

	if len(notFound) > 0 {
		notFoundStrings, err := utils.ConvertIntsToStrings(notFound)
		if err != nil {
			return nil, err
		}

		newCompanies, err := l.GetCompanies(ctx, notFoundStrings)

		if err != nil {
			return nil, err
		}

		err = r.Cache.AddCompanies(ctx, newCompanies)

		if err != nil {
			return nil, err
		}

		for _, company := range newCompanies {
			idInt, err := utils.StringToInt(&company.ID)

			if err != nil {
				return nil, err
			}

			companies[idInt] = company
		}
	}

	// convert companies to []*model.Company based on ids order
	companiesList := []*model.Company{}
	for _, id := range intIds {
		companiesList = append(companiesList, companies[id])
	}

	return companiesList, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
