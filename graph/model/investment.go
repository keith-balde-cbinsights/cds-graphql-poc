package model

import "time"

type Investment struct {
	ID            string    `json:"id"`
	RoundName     string    `json:"roundName"`
	Date          time.Time `json:"date"`
	Amount        float64   `json:"amount"`
	Valuation     float64   `json:"valuation"`
	ReceiverIDOrg int       `json:"receiverID"`
	Receiver      *Company  `json:"receiver"`
	Investor      *Company  `json:"investor"`
}
