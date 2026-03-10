package service

import "github.com/dwisarut/dealmaxxingCLI/internal/model"

func MakeRedirectLink(lists []model.SearchGameID) []model.SearchGameID {
	for i := range lists {
		lists[i].Redirect = "https://www.cheapshark.com/redirect?dealID=" + lists[i].DealID
	}
	return lists
}
