package data

type GameData struct {
	internalName       string
	title              string
	dealID             string
	isOnSale           string
	saving             string
	salePrice          string
	metacriticScore    string
	steamRatingPercent string
}

type InputData struct {
	suffix  string
	command string
	param   string
	name    string
}
