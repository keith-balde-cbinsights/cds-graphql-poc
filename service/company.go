package service

import (
	"cds-graphql-poc/client/profileservice"
	"cds-graphql-poc/dto"
	"cds-graphql-poc/graph/model"
	"cds-graphql-poc/graph/utils"
	"context"
	"fmt"
	"strconv"
)

type CompanyService interface {
	GetCompaniesById(ctx context.Context, ids []*string) ([]*model.Company, error)
}

type companyService struct {
	profileServiceClient profileservice.Client
}

func NewCompanyService() *companyService {
	return &companyService{
		profileServiceClient: profileservice.NewClient(),
	}
}

func (s *companyService) GetCompaniesById(ctx context.Context, ids []*string) ([]*model.Company, error) {
	intIds := []int{}
	for _, id := range ids {
		newId, err := strconv.Atoi(*id)

		if err != nil {
			return nil, fmt.Errorf("error converting id to int: %v", err)
		}

		intIds = append(intIds, newId)
	}

	profiles, err := s.profileServiceClient.GetProfilesById(ctx, intIds)

	if err != nil {
		return nil, fmt.Errorf("error getting profiles by id: %v", err)
	}

	companies := []*model.Company{}

	for _, profile := range profiles {

		profileURL, err := utils.GenerateProfileURL(profile.IdCbiEntity)

		if err != nil {
			return nil, fmt.Errorf("error generating profile URL: %v", err)
		}

		location := &model.CompanyLocation{}
		if profile.Address != nil {
			location = &model.CompanyLocation{
				Country: profile.Address.Country,
				State:   profile.Address.State,
				City:    profile.Address.City,
			}

		}

		companies = append(companies, &model.Company{
			ID:         strconv.Itoa(profile.IdCbiEntity),
			OrgID:      int32(profile.IdCbiEntity),
			InvestorID: int32(profile.IdInvestor),
			CompanyID:  int32(profile.IdCompany),
			Name:       profile.Name,
			Status:     profile.Status,
			Website:    profile.Url,
			ProfileURL: profileURL,
			Location:   location,
		})
	}

	return companies, nil
}

func (s *companyService) GetSummaryKPIForCompanies(ctx context.Context, ids []string) ([]*dto.KPISummary, error) {
	intIds := []int{}
	for _, id := range ids {
		newId, err := strconv.Atoi(id)

		if err != nil {
			return nil, fmt.Errorf("error converting id to int: %v", err)
		}

		intIds = append(intIds, newId)
	}

	summaryKPIs, err := s.profileServiceClient.GetSummaryKPIForCompanies(ctx, intIds)

	if err != nil {
		return nil, fmt.Errorf("error getting summary KPIs for companies: %v", err)
	}

	summaryKPIsDTO := []*dto.KPISummary{}

	for _, summaryKPI := range summaryKPIs {
		summaryKPIsDTO = append(summaryKPIsDTO, &dto.KPISummary{
			MarketCap: summaryKPI.MarketCap,
		})
	}

	return summaryKPIsDTO, nil
}
