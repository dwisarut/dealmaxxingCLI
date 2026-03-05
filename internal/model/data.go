package model

type GameData struct {
	InternalName       string `json:"internalName"`
	Title              string `json:"title"`
	DealID             string `json:"dealID"`
	IsOnSale           string `json:"isOnSale"`
	Saving             string `json:"saving"`
	SalePrice          string `json:"salePrice"`
	MetacriticScore    string `json:"metacriticScore"`
	SteamRatingPercent string `json:"steamRatingPercent"`
}

type InputDataCLI struct {
	Search   string
	Query    string
	Top      int
	Command  bool
	MaxPrice int
}
