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

type SearchGameID struct {
	Title     string `json:"title"`
	GameIDTag string `json:"gameID"`
	DealID    string `json:"dealID"`
	SalePrice string `json:"salePrice"`
	StoreID   string `json:"storeID"`
	Redirect  string
}

type QueryParseType struct {
	Command        string
	SortBy         string
	Limit          string
	RatingSource   string
	MinRatingScore string
}
