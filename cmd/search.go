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
	var pageNum int = 1

	fmt.Println()
	fmt.Println("Searching...")
	fmt.Println()

	for {
		var lists []model.SearchGameID = api.GetDealFromTitle(title, pageNum, 5)
		var displayLists []model.SearchGameID = service.MakeRedirectLink(lists)

		fmt.Println(color.HiMagentaString("Page:"), pageNum)
		fmt.Println()
		for _, list := range displayLists {
			fmt.Println(color.HiCyanString(list.Title))
			fmt.Println("ID:", list.GameIDTag)
			fmt.Println("Prices:", list.SalePrice, "$")
			fmt.Println(color.GreenString("Link:"), color.GreenString(list.Redirect))
			fmt.Println()
		}
		fmt.Println("Commands:", color.HiYellowString("prev (p)"), "|", color.HiGreenString("next (n)"), "|", color.HiRedString("cancel (cc)"))
		fmt.Print("> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "n":
			if pageNum < len(lists) {
				pageNum++
			}
			fmt.Println()

		case "p":
			if pageNum > 1 {
				pageNum--
			}
			fmt.Println()

		case "cc":
			println("Exiting search function!")
			fmt.Println()
			return

		default:
			println("Unknown command.")
		}
	}
}
