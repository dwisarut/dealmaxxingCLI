package data

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

type InputData struct {
	Suffix  string `json:"suffix"`
	Command string `json:"command"`
	Param   string `json:"param"`
	Name    string `json:"name"`
}
