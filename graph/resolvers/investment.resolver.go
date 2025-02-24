package resolvers

import (
	"cds-graphql-poc/graph"
	"cds-graphql-poc/graph/model"
	"context"
	"fmt"
)

func (r *investmentResolver) Receiver(ctx context.Context, obj *model.Investment) (*model.Company, error) {
	if company, exists := r.Cache.GetCompany(ctx, int(obj.ReceiverIDOrg)); exists {
		return company, nil
	} else {
		id := fmt.Sprintf("%d", obj.ReceiverIDOrg)

		company, err := r.Loaders.GetCompany(ctx, &id)
		if err != nil {
			return nil, err
		}

		r.Cache.AddCompany(ctx, company)

		return company, nil
	}
}

func (r *Resolver) Investment() graph.InvestmentResolver { return &investmentResolver{r} }

type investmentResolver struct{ *Resolver }
