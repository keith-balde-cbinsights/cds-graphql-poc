package contextcache

import (
	"context"
	"fmt"
	"strconv"

	"cds-graphql-poc/dto"
	"cds-graphql-poc/graph/model"
)

type CompanyCacheKey struct{}
type InvestmentCacheKey struct{}
type SummaryKPICacheKey struct{}

// Context keys for caching
func GetCompanyCacheKey() *CompanyCacheKey {
	return &CompanyCacheKey{}
}

func GetInvestmentCacheKey() *InvestmentCacheKey {
	return &InvestmentCacheKey{}
}

func GetSummaryKPICacheKey() *SummaryKPICacheKey {
	return &SummaryKPICacheKey{}
}

type Cache struct{}

func NewCache() *Cache {
	return &Cache{}
}

// GetCompanyCache retrieves the company cache from the context
func GetCompanyCache(ctx context.Context) map[int]*model.Company {
	cache, _ := ctx.Value(CompanyCacheKey{}).(map[int]*model.Company)
	return cache
}

// GetInvestmentCache retrieves the investment cache from the context
func GetInvestmentCache(ctx context.Context) map[int][]*model.Investment {
	cache, _ := ctx.Value(InvestmentCacheKey{}).(map[int][]*model.Investment)
	return cache
}

// GetSummaryKPICache retrieves the summary KPI cache from the context
func GetSummaryKPICache(ctx context.Context) map[int]*dto.KPISummary {
	cache, _ := ctx.Value(SummaryKPICacheKey{}).(map[int]*dto.KPISummary)
	return cache
}

// WithCompanyCache attaches a cache map to the context
func WithCompanyCache(ctx context.Context) context.Context {
	return context.WithValue(ctx, CompanyCacheKey{}, make(map[int]*model.Company))
}

// WithInvestmentCache attaches a cache map to the context
func WithInvestmentCache(ctx context.Context) context.Context {
	return context.WithValue(ctx, InvestmentCacheKey{}, make(map[int][]*model.Investment))
}

// WithSummaryKPICache attaches a cache map to the context
func WithSummaryKPICache(ctx context.Context) context.Context {
	return context.WithValue(ctx, SummaryKPICacheKey{}, make(map[int]*dto.KPISummary))
}

func (c *Cache) GetCompany(ctx context.Context, id int) (*model.Company, bool) {
	cache := GetCompanyCache(ctx)
	company, ok := cache[id]
	return company, ok
}

func (c *Cache) GetCompanies(ctx context.Context, ids []int) (map[int]*model.Company, []int) {
	cache := GetCompanyCache(ctx)
	companies := map[int]*model.Company{}
	notFound := []int{}

	for _, id := range ids {
		if company, ok := cache[id]; ok {
			companies[id] = company
		} else {
			notFound = append(notFound, id)
		}
	}

	return companies, notFound
}

func (c *Cache) GetInvestmentsForCompany(ctx context.Context, id int) ([]*model.Investment, bool) {
	cache := GetInvestmentCache(ctx)
	investments, ok := cache[id]
	return investments, ok
}

func (c *Cache) GetSummaryKPI(ctx context.Context, id int) (*dto.KPISummary, bool) {
	cache := GetSummaryKPICache(ctx)
	summaryKPI, ok := cache[id]
	return summaryKPI, ok
}

func (c *Cache) AddCompany(ctx context.Context, company *model.Company) error {
	cache := GetCompanyCache(ctx)

	id, err := strconv.Atoi(company.ID)
	if err != nil {
		return fmt.Errorf("failed to convert company.ID of %s to int when adding company to cache: %w", company.ID, err)
	}
	cache[id] = company

	return nil
}

func (c *Cache) AddCompanies(ctx context.Context, companies []*model.Company) error {
	for _, company := range companies {
		err := c.AddCompany(ctx, company)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Cache) AddInvestmentsForCompany(ctx context.Context, idOrg string, investments []*model.Investment) error {
	cache := GetInvestmentCache(ctx)

	id, err := strconv.Atoi(idOrg)
	if err != nil {
		return fmt.Errorf("failed to convert idOrg of %s to int when adding investments to cache: %w", idOrg, err)
	}
	cache[id] = investments

	return nil
}

func (c *Cache) AddSummaryKPI(ctx context.Context, idOrg string, summary *dto.KPISummary) error {
	cache := GetSummaryKPICache(ctx)

	id, err := strconv.Atoi(idOrg)
	if err != nil {
		return fmt.Errorf("failed to convert idOrg of %s to int when adding KPI summary to cache: %w", idOrg, err)
	}
	cache[id] = summary

	return nil
}
