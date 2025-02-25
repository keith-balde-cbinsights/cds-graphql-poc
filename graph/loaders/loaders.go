package loaders

import (
	"cds-graphql-poc/dto"
	"cds-graphql-poc/graph/model"
	"cds-graphql-poc/service/company"
	"context"

	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	CompanyLoader    *dataloadgen.Loader[*string, *model.Company]
	SummaryKPILoader *dataloadgen.Loader[int, *dto.KPISummary]
	InvestmentLoader *dataloadgen.Loader[int, []*model.Investment]
}

func GetLoadersKey() ctxKey {
	return loadersKey
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders() *Loaders {
	companyService := company.NewService()

	return &Loaders{
		CompanyLoader:    dataloadgen.NewLoader(companyService.GetCompaniesById),
		SummaryKPILoader: dataloadgen.NewLoader(companyService.GetSummaryKPIForCompanies),
		InvestmentLoader: dataloadgen.NewLoader(companyService.GetInvestments),
	}
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
