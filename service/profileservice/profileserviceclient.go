package profileservice

import (
	r "cds-graphql-poc/example/response"
	"context"
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
}

type client struct {
}

func NewClient() *client {
	r.PopulateProfilesMap()

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
	profiles := []*r.GetProfileResponse{}

	for _, id := range ids {
		profiles = append(profiles, r.Profiles[id])
	}

	return profiles, nil
}
