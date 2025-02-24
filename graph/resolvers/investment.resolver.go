package resolvers

import (
	"cds-graphql-poc/graph"
	"cds-graphql-poc/graph/model"
	"context"
	"fmt"
)

func (r *investmentResolver) Receiver(ctx context.Context, obj *model.Investment) (*model.Company, error) {
	id := fmt.Sprintf("%d", obj.ReceiverIDOrg)
	return r.Loaders.GetCompany(ctx, &id)
}

func (r *Resolver) Investment() graph.InvestmentResolver { return &investmentResolver{r} }

type investmentResolver struct{ *Resolver }
