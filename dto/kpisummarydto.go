package dto

type KPISummary struct {
	MarketCap float64
	Ceo       *KeyPerson
}

type KeyPerson struct {
	Id       int
	FullName string
	Title    string
}
