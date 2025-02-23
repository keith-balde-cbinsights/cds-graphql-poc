package response

import "time"

type GetProfileResponse struct {
	IdCbiEntity int
	IdCompany   int
	IdInvestor  int
	Name        string
	Url         string
	Status      string
	Address     *Address
}

type GetSummaryKPIsResponse struct {
	IdCompany    int
	Ceo          *KeyPerson
	MarketCap    float64
	TotalFunding float64
}

type GetInvestmentsResponse struct {
	Investments []*Investment
}

type Investment struct {
	Id            int
	RoundName     string
	Date          time.Time
	Amount        float64
	Valuation     float64
	InvestedIdOrg int
}

type KeyPerson struct {
	Id       int
	FullName string
	Title    string
}

type Address struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zip     string
	Country string
}

var Profiles = map[int]*GetProfileResponse{}
var SummaryKPIs = map[int]*GetSummaryKPIsResponse{}
var Investments = map[int]*GetInvestmentsResponse{}

func PopulateProfilesMap() {
	Profiles[1] = &GetProfileResponse{
		IdCompany:   1,
		IdCbiEntity: 1,
		IdInvestor:  1,
		Name:        "CB Insights",
		Url:         "https://www.cbinsights.com",
		Status:      "ACTIVE",
		Address: &Address{
			Street1: "498 7th Ave",
			City:    "New York",
			State:   "NY",
			Zip:     "10018",
			Country: "USA",
		},
	}
	Profiles[2] = &GetProfileResponse{
		IdCompany:   2,
		IdInvestor:  2,
		IdCbiEntity: 2,
		Name:        "Google",
		Url:         "https://www.google.com",
		Status:      "ACTIVE",
		Address: &Address{
			Street1: "1600 Amphitheatre Parkway",
			City:    "Mountain View",
			State:   "CA",
			Zip:     "94043",
			Country: "USA",
		},
	}
	Profiles[3] = &GetProfileResponse{
		IdCompany:   3,
		IdInvestor:  3,
		IdCbiEntity: 3,
		Name:        "Meta",
		Url:         "https://www.meta.com",
		Status:      "ACTIVE",
		Address: &Address{
			Street1: "1 Hacker Way",
			City:    "Menlo Park",
			State:   "CA",
			Zip:     "94025",
			Country: "USA",
		},
	}
}

func PopulateSummaryKPIs() {
	SummaryKPIs[1] = &GetSummaryKPIsResponse{
		MarketCap:    1000,
		TotalFunding: 9999,
		Ceo: &KeyPerson{
			Id:       1,
			FullName: "Manlio Carrelli",
			Title:    "Chief Executive Officer",
		},
	}
	SummaryKPIs[2] = &GetSummaryKPIsResponse{
		MarketCap:    1500,
		TotalFunding: 8888,
		Ceo: &KeyPerson{
			Id:       2,
			FullName: "Sundar Pichai",
			Title:    "CEO",
		},
	}
	SummaryKPIs[3] = &GetSummaryKPIsResponse{
		MarketCap:    800,
		TotalFunding: 5555,
		Ceo: &KeyPerson{
			Id:       3,
			FullName: "Mark Zuckerberg",
			Title:    "Chief Executive Officer",
		},
	}
}

func PopulateInvestments() {
	Investments[1] = &GetInvestmentsResponse{
		Investments: []*Investment{
			{
				Id:            1,
				RoundName:     "Series A",
				Date:          time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Amount:        1000000,
				Valuation:     10000000,
				InvestedIdOrg: 1,
			},
			{
				Id:            2,
				RoundName:     "Series B",
				Date:          time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				Amount:        2000000,
				Valuation:     20000000,
				InvestedIdOrg: 1,
			},
		},
	}
	Investments[2] = &GetInvestmentsResponse{
		Investments: []*Investment{
			{
				Id:            3,
				RoundName:     "Series C",
				Date:          time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				Amount:        3000000,
				Valuation:     30000000,
				InvestedIdOrg: 2,
			},
		},
	}
	Investments[3] = &GetInvestmentsResponse{
		Investments: []*Investment{
			{
				Id:            4,
				RoundName:     "Series D",
				Date:          time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Amount:        4000000,
				Valuation:     40000000,
				InvestedIdOrg: 3,
			},
		},
	}
}
