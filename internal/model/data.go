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
	Title     string `json:"external"`
	GameIDTag string `json:"gameID"`
	DealID    string `json:"cheapestDealID"`
	SalePrice string `json:"cheapest"`
	Redirect  string
	StoreName string
}

type StoreLists struct {
	ID   string `json:"storeID"`
	Name string `json:"storeName"`
}

type GetGameID struct {
	Info  InfoAPI   `json:"info"`
	Deals []DealAPI `json:"deals"`
}

type InfoAPI struct {
	Title string `json:"title"`
}

type DealAPI struct {
	StoreID   string `json:"storeID"`
	DealID    string `json:"dealID"`
	Price     string `json:"price"`
	Savings   string `json:"savings"`
	StoreName string
	Redirect  string
}

type GameIdentifier struct {
	Name   string `json:"external"`
	GameID string `json:"gameID"`
}
