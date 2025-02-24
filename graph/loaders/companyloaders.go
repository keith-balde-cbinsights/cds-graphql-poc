package loaders

import (
	"cds-graphql-poc/dto"
	"cds-graphql-poc/graph/model"
	"context"
	"fmt"
)

func (l *Loaders) GetCompany(ctx context.Context, id *string) (*model.Company, error) {
	fmt.Println("Using loader GetCompany", id)

	return l.CompanyLoader.Load(ctx, id)
}

func (l *Loaders) GetCompanies(ctx context.Context, ids []*string) ([]*model.Company, error) {
	fmt.Println("Using loader GetCompanies", ids)

	return l.CompanyLoader.LoadAll(ctx, ids)
}

func (l *Loaders) GetSummaryKPI(ctx context.Context, id int) (*dto.KPISummary, error) {
	fmt.Println("Using loader GetSummaryKPI", id)

	return l.SummaryKPILoader.Load(ctx, id)
}

func (l *Loaders) GetSummaryKPIs(ctx context.Context, ids []int) ([]*dto.KPISummary, error) {
	fmt.Println("Using loader GetSummaryKPIs", ids)

	return l.SummaryKPILoader.LoadAll(ctx, ids)
}

func (l *Loaders) GetInvestmentsForCompany(ctx context.Context, id int) ([]*model.Investment, error) {
	fmt.Println("Using loader GetInvestmentsForCompany", id)

	return l.InvestmentLoader.Load(ctx, id)
}

func (l *Loaders) GetInvestmentsForCompanies(ctx context.Context, ids []int) ([][]*model.Investment, error) {
	fmt.Println("Using loader GetInvestmentsForCompanies", ids)

	return l.InvestmentLoader.LoadAll(ctx, ids)
}
