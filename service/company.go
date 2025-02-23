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
	GetCompaniesById(ctx context.Context, ids []*string) ([]*model.Company, []error)
	GetSummaryKPIForCompanies(ctx context.Context, ids []int) ([]*dto.KPISummary, []error)
	GetInvestments(ctx context.Context, ids []int) ([][]*model.Investment, []error)
}

type companyService struct {
	profileServiceClient profileservice.Client
}

func NewCompanyService() *companyService {
	return &companyService{
		profileServiceClient: profileservice.NewClient(),
	}
}

func (s *companyService) GetCompaniesById(ctx context.Context, ids []*string) ([]*model.Company, []error) {
	intIds := []int{}
	for _, id := range ids {
		newId, err := strconv.Atoi(*id)

		if err != nil {
			return nil, []error{fmt.Errorf("error converting id to int: %v", err)}
		}

		intIds = append(intIds, newId)
	}

	profiles, err := s.profileServiceClient.GetProfilesById(ctx, intIds)

	if err != nil {
		return nil, []error{fmt.Errorf("error getting profiles by id: %v", err)}
	}

	companies := []*model.Company{}

	for _, profile := range profiles {

		profileURL, err := utils.GenerateProfileURL(profile.IdCbiEntity)

		if err != nil {
			return nil, []error{fmt.Errorf("error generating profile URL: %v", err)}
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

func (s *companyService) GetSummaryKPIForCompanies(ctx context.Context, ids []int) ([]*dto.KPISummary, []error) {
	summaryKPIs, err := s.profileServiceClient.GetSummaryKPIForCompanies(ctx, ids)

	if err != nil {
		return nil, []error{fmt.Errorf("error getting summary KPIs for companies: %v", err)}
	}

	summaryKPIsDTO := []*dto.KPISummary{}

	for _, summaryKPI := range summaryKPIs {
		ceo := &dto.KeyPerson{
			Id:       summaryKPI.Ceo.Id,
			FullName: summaryKPI.Ceo.FullName,
			Title:    summaryKPI.Ceo.Title,
		}

		summaryKPIsDTO = append(summaryKPIsDTO, &dto.KPISummary{
			MarketCap:    summaryKPI.MarketCap,
			TotalFunding: summaryKPI.TotalFunding,
			Ceo:          ceo,
		})
	}

	return summaryKPIsDTO, nil
}

func (s *companyService) GetInvestments(ctx context.Context, ids []int) ([][]*model.Investment, []error) {
	companies, err := s.profileServiceClient.GetInvestments(ctx, ids)

	if err != nil {
		return nil, []error{fmt.Errorf("error getting investments: %v", err)}
	}

	result := [][]*model.Investment{}

	for _, company := range companies {
		resultItem := []*model.Investment{}
		for _, investment := range company.Investments {
			resultItem = append(resultItem, &model.Investment{
				ID:        strconv.Itoa(investment.Id),
				RoundName: investment.RoundName,
				Date:      investment.Date,
				Amount:    investment.Amount,
				Valuation: investment.Valuation,
			})
		}

		result = append(result, resultItem)
	}

	return result, nil
}
