package cmd

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/dwisarut/dealmaxxingCLI/internal/service"
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

		fmt.Println("Page", pageNum)
		fmt.Println()
		for _, list := range displayLists {
			fmt.Println("Title:", list.Title)
			fmt.Println("ID:", list.GameIDTag)
			fmt.Println("Prices:", list.SalePrice, "$")
			fmt.Println("Link:", list.Redirect)
			fmt.Println()
		}
		fmt.Println("Commands: prev | next | cc (cancel)")
		fmt.Print("> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "next":
			pageNum++
			fmt.Println()

		case "prev":
			if pageNum > 1 {
				pageNum--
			}
			fmt.Println()

		case "cc":
			return

		default:
			println("Unknown command.")
		}
	}
}
