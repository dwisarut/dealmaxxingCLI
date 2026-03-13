package service

import (
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
)

func MakeSearchRedirect(lists []model.SearchGameID) []model.SearchGameID {
	for i := range lists {
		lists[i].Redirect = "https://www.cheapshark.com/redirect?dealID=" + lists[i].DealID
	}
	return lists
}

func MatchSearchStore(lists []model.SearchGameID, storeLookup map[string]string) []model.SearchGameID {
	for i := range lists {
		if val, ok := storeLookup[lists[i].StoreID]; ok { // JS Map.has() & Map.get() equivalent
			lists[i].StoreName = val
		}
	}

	return lists
}

func IndexingStore(stores []model.StoreLists) map[string]string {
	var storeLookup map[string]string
	storeLookup = make(map[string]string)

	for _, store := range stores {
		storeLookup[store.ID] = store.Name
	}

	return storeLookup
}

func MakeGetRedirect(list model.GetGameID) model.GetGameID {
	for i := range list.Deals {
		list.Deals[i].Redirect = "https://www.cheapshark.com/redirect?dealID=" + list.Deals[i].DealID
	}

	return list
}

func MatchGetStore(list model.GetGameID, storeLookup map[string]string) model.GetGameID {
	for i := range list.Deals {
		if val, ok := storeLookup[list.Deals[i].StoreID]; ok { // JS Map.has() & Map.get() equivalent
			list.Deals[i].StoreName = val
		}
	}

	return list
}
