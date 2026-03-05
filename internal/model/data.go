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

type SearchParseType struct {
	Command string
	Arg     string
}

type QueryParseType struct {
	Command        string
	SortBy         string
	Limit          string
	RatingSource   string
	MinRatingScore string
}
