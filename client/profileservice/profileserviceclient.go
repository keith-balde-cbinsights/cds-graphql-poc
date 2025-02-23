package profileservice

import (
	r "cds-graphql-poc/example/response"
	"context"
	"fmt"
)

type Client interface {
	GetProfile(
		ctx context.Context,
		id int,
	) (*r.GetProfileResponse, error)
	GetProfilesById(
		ctx context.Context,
		ids []int,
	) ([]*r.GetProfileResponse, error)
	GetSummaryKPIForCompanies(
		ctx context.Context,
		ids []int,
	) ([]*r.GetSummaryKPIsResponse, error)
}

type client struct {
}

func NewClient() *client {
	r.PopulateProfilesMap()
	r.PopulateSummaryKPIs()

	return &client{}
}

func (c *client) GetProfile(
	ctx context.Context,
	id int,
) (*r.GetProfileResponse, error) {
	return r.Profiles[id], nil
}

func (c *client) GetProfilesById(
	ctx context.Context,
	ids []int,
) ([]*r.GetProfileResponse, error) {
	fmt.Printf("Fetching profiles by ids: %v\n", ids)

	profiles := []*r.GetProfileResponse{}

	for _, id := range ids {
		profiles = append(profiles, r.Profiles[id])
	}

	return profiles, nil
}

func (c *client) GetSummaryKPIForCompanies(
	ctx context.Context,
	ids []int,
) ([]*r.GetSummaryKPIsResponse, error) {
	fmt.Printf("Fetching summary KPIs for companies: %v\n", ids)

	summaryKPIs := []*r.GetSummaryKPIsResponse{}

	for _, id := range ids {
		summaryKPIs = append(summaryKPIs, r.SummaryKPIs[id])
	}

	return summaryKPIs, nil
}
