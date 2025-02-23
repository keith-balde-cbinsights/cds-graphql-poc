package loaders

import (
	"cds-graphql-poc/dto"
	"cds-graphql-poc/graph/model"
	"context"
	"fmt"
)

func GetCompany(ctx context.Context, id *string) (*model.Company, error) {
	fmt.Println("Using loader GetCompany", id)
	loaders := For(ctx)

	return loaders.CompanyLoader.Load(ctx, id)
}

func GetCompanies(ctx context.Context, ids []*string) ([]*model.Company, error) {
	fmt.Println("Using loader GetCompanies", ids)
	loaders := For(ctx)

	return loaders.CompanyLoader.LoadAll(ctx, ids)
}

func GetSummaryKPI(ctx context.Context, id int) (*dto.KPISummary, error) {
	fmt.Println("Using loader GetSummaryKPI", id)
	loaders := For(ctx)

	return loaders.SummaryKPILoader.Load(ctx, id)
}

func GetSummaryKPIs(ctx context.Context, ids []int) ([]*dto.KPISummary, error) {
	fmt.Println("Using loader GetSummaryKPIs", ids)
	loaders := For(ctx)

	return loaders.SummaryKPILoader.LoadAll(ctx, ids)
}

func GetInvestmentsForCompany(ctx context.Context, id int) ([]*model.Investment, error) {
	fmt.Println("Using loader GetInvestmentsForCompany", id)
	loaders := For(ctx)

	return loaders.InvestmentLoader.Load(ctx, id)
}

func GetInvestmentsForCompanies(ctx context.Context, ids []int) ([][]*model.Investment, error) {
	fmt.Println("Using loader GetInvestmentsForCompanies", ids)
	loaders := For(ctx)

	return loaders.InvestmentLoader.LoadAll(ctx, ids)
}
