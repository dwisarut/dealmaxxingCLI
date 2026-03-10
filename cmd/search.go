package cmd

import (
	"fmt"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/dwisarut/dealmaxxingCLI/internal/service"
)

func SearchHandler(input string) {
	var title string = CommonParser(input)
	fmt.Println("Searching...")
	fmt.Println()
	var lists []model.SearchGameID = api.GetDealId(title)
	var displayLists []model.SearchGameID = service.MakeRedirectLink(lists)

	for _, list := range displayLists {
		fmt.Println("Title:", list.Title)
		fmt.Println("ID:", list.GameIDTag)
		fmt.Println("Prices:", list.SalePrice, "$")
		fmt.Println("Link:", list.Redirect)
		fmt.Println()
	}
}
