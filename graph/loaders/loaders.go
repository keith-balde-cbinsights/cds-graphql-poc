package loaders

import (
	"cds-graphql-poc/dto"
	"cds-graphql-poc/graph/model"
	"cds-graphql-poc/service/company"
	"context"
	"net/http"

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

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(companyService company.Service) *Loaders {
	return &Loaders{
		CompanyLoader:    dataloadgen.NewLoader(companyService.GetCompaniesById),
		SummaryKPILoader: dataloadgen.NewLoader(companyService.GetSummaryKPIForCompanies),
		InvestmentLoader: dataloadgen.NewLoader(companyService.GetInvestments),
	}
}

// Middleware injects data loaders into the context
func Middleware(
	companyService company.Service,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := NewLoaders(companyService)

		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loaders))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
