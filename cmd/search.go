package cmd

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/dwisarut/dealmaxxingCLI/internal/service"
	"github.com/fatih/color"
)

func SearchHandler(reader *bufio.Reader, input string) {
	var title string = CommonParser(input)

	if title == "" {
		return
	}

	fmt.Println()
	color.Green("Searching...")
	fmt.Println()

	var lists []model.SearchGameID = api.GetDealFromTitle(title)

	if len(lists) == 0 || len(lists) == 60 {
		fmt.Println(color.HiRedString("No games founded."))
		fmt.Println("Exiting...")
		fmt.Println()
		return
	}

	var displayLists []model.SearchGameID = service.MakeSearchRedirect(lists)

	fmt.Println(color.HiWhiteString("Total results:"), len(lists))
	fmt.Println()

	for _, list := range displayLists {
		fmt.Println(color.HiCyanString(list.Title))
		fmt.Println(color.HiMagentaString("ID:"), color.MagentaString(list.GameIDTag))
		fmt.Println(color.HiYellowString("Prices:"), color.YellowString(list.SalePrice), color.YellowString("$"))
		fmt.Println(color.HiWhiteString("Link:"), color.GreenString(list.Redirect))
		fmt.Println()
	}

	fmt.Println(strings.Repeat("_", 100))
}
