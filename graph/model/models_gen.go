// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Company struct {
	ID            string           `json:"id"`
	OrgID         int32            `json:"orgID"`
	InvestorID    int32            `json:"investorID"`
	CompanyID     int32            `json:"companyID"`
	Name          string           `json:"name"`
	Website       string           `json:"website"`
	ProfileURL    string           `json:"profileUrl"`
	Status        string           `json:"status"`
	Location      *CompanyLocation `json:"location"`
	FundingRounds []*FundingRound  `json:"fundingRounds"`
	Investments   []*Investment    `json:"investments"`
	MarketCap     float64          `json:"marketCap"`
	TotalRaised   float64          `json:"totalRaised"`
	Ceo           *KeyPerson       `json:"ceo,omitempty"`
}

type CompanyLocation struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type FundingRound struct {
	ID             string     `json:"id"`
	RoundName      string     `json:"roundName"`
	Date           time.Time  `json:"date"`
	Amount         float64    `json:"amount"`
	Valuation      float64    `json:"valuation"`
	Receiver       *Company   `json:"receiver"`
	LeadInvestor   *Company   `json:"leadInvestor,omitempty"`
	OtherInvestors []*Company `json:"otherInvestors"`
}

type KeyPerson struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Company *Company `json:"company"`
}

type Mutation struct {
}

type NewCompany struct {
	Name  string `json:"name"`
	OrgID string `json:"orgId"`
}

type Query struct {
}

type Schema struct {
	Query    *Query    `json:"query,omitempty"`
	Mutation *Mutation `json:"mutation,omitempty"`
}

type Summary struct {
	Headline string    `json:"headline"`
	Basics   string    `json:"basics"`
	Details  []*string `json:"details"`
	Sources  []*string `json:"sources"`
}
