package loaders

import (
	"cds-graphql-poc/dto"
	"cds-graphql-poc/graph/model"
	"cds-graphql-poc/utils"
	"context"
	"fmt"
)

func GetCompany(ctx context.Context, id *string) (*model.Company, error) {
	loaders := For(ctx)

	idInt, err := utils.StringToInt(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to int: %w", err)
	}

	fmt.Println("Using loader GetCompany", idInt)

	return loaders.CompanyLoader.Load(ctx, id)
}

func GetCompanies(ctx context.Context, ids []*string) ([]*model.Company, error) {
	loaders := For(ctx)

	intIds, err := utils.ConvertStringsToInts(ids)

	if err != nil {
		return nil, fmt.Errorf("failed to convert ids to ints: %w", err)
	}

	fmt.Println("Using loader GetCompanies", intIds)

	return loaders.CompanyLoader.LoadAll(ctx, ids)
}

func GetSummaryKPI(ctx context.Context, id int) (*dto.KPISummary, error) {
	loaders := For(ctx)

	fmt.Println("Using loader GetSummaryKPI", id)

	return loaders.SummaryKPILoader.Load(ctx, id)
}

func GetSummaryKPIs(ctx context.Context, ids []int) ([]*dto.KPISummary, error) {
	loaders := For(ctx)

	fmt.Println("Using loader GetSummaryKPIs", ids)

	return loaders.SummaryKPILoader.LoadAll(ctx, ids)
}

func GetInvestmentsForCompany(ctx context.Context, id int) ([]*model.Investment, error) {
	loaders := For(ctx)

	fmt.Println("Using loader GetInvestmentsForCompany", id)

	return loaders.InvestmentLoader.Load(ctx, id)
}

func GetInvestmentsForCompanies(ctx context.Context, ids []int) ([][]*model.Investment, error) {
	loaders := For(ctx)

	fmt.Println("Using loader GetInvestmentsForCompanies", ids)

	return loaders.InvestmentLoader.LoadAll(ctx, ids)
}
