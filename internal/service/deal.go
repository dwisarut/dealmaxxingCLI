package service

import (
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
)

func MakeRedirectLink(lists []model.SearchGameID) []model.SearchGameID {
	for i := range lists {
		lists[i].Redirect = "https://www.cheapshark.com/redirect?dealID=" + lists[i].DealID
	}
	return lists
}

func MatchStore(lists []model.SearchGameID, stores []model.StoreLists) []model.SearchGameID {
	storeLookup := make(map[string]string) // Hashmap

	for _, store := range stores {
		storeLookup[store.ID] = store.Name
	}

	for i := range lists {
		if val, ok := storeLookup[lists[i].StoreID]; ok { // JS Map.has() & Map.get() equivalent
			lists[i].StoreName = val
		}
	}

	return lists
}
