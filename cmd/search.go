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

func SearchHandler(reader *bufio.Reader, input string, storeIndex map[string]string) {
	var title string = CommonParser(input)
	var pageNum int = 1
	var pageSize int = 5

	fmt.Println()
	color.Green("Searching...")
	fmt.Println()

	for {
		var lists []model.SearchGameID = api.GetDealFromTitle(title, pageNum, pageSize)

		if len(lists) == 0 {
			fmt.Println(color.HiRedString("No games founded."))
			fmt.Println("Exiting...")
			fmt.Println()
			return
		}

		var displayLists []model.SearchGameID = service.MakeSearchRedirect(lists)
		displayLists = service.MatchSearchStore(displayLists, storeIndex)

		fmt.Println(color.HiWhiteString("Page:"), pageNum, len(lists))
		fmt.Println()

		for _, list := range displayLists {
			fmt.Println(color.HiCyanString(list.Title))
			fmt.Println(color.HiBlueString("Store:"), color.BlueString(list.StoreName))
			fmt.Println(color.HiMagentaString("ID:"), color.MagentaString(list.GameIDTag))
			fmt.Println(color.HiYellowString("Prices:"), color.YellowString(list.SalePrice), color.YellowString("$"))
			fmt.Println(color.HiWhiteString("Link:"), color.GreenString(list.Redirect))
			fmt.Println()
		}

		fmt.Println("Commands:", color.HiYellowString("prev (p)"), "|", color.HiGreenString("next (n)"), "|", color.HiRedString("cancel (cc)"))
		fmt.Print("> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "n":
			pageNum++
			fmt.Println()

		case "p":
			if pageNum > 1 {
				pageNum--
			}
			fmt.Println()

		case "cc":
			println(color.HiRedString("Exiting search function!"))
			fmt.Println()
			return

		default:
			println("Unknown command.")
		}
	}
}
