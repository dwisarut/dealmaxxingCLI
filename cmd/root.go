package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/dwisarut/dealmaxxingCLI/internal/service"
)

func InitCLI() {
	reader := bufio.NewReader(os.Stdin)

	initMessage()

	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		case strings.HasPrefix(input, "cmd"):
			showCommand()

		case strings.HasPrefix(input, "query"):
			QueryParser(input)
			fmt.Println("Querying...")

		case strings.HasPrefix(input, "search"):
			var title string = CommonParser(input)
			fmt.Println("Searching...")
			fmt.Println()
			var lists []model.SearchGameID = api.GetDealId(title)
			var displayLists []model.SearchGameID = service.MakeRedirectLink(lists)

			for _, list := range displayLists {
				fmt.Println("Title:", list.Title)
				fmt.Println("Prices:", list.SalePrice, "$")
				fmt.Println("Link:", list.Redirect)
				fmt.Println()
			}

		case strings.HasPrefix(input, "get"):
			var title string = CommonParser(input)
			fmt.Println("Getting...")
			api.GetDealId(title)

		case input == "exit":
			fmt.Println("Exiting DealmaxxingCLI...")
			return

		default:
			fmt.Println("Unknown command")
		}
	}

}

func showCommand() {
	fmt.Println("Here's an available commands:")
	fmt.Printf("%-10s %s\n", "search", "Searching for a specific game")
	fmt.Printf("%-10s %s\n", "query", "Query lists of game with specific parameter")
}

func initMessage() {
	fmt.Println("DealmaxxingCLI")
	fmt.Println("Your trusty tools for finding cheap game!")
	fmt.Println("Type cmd to find an available commands.")
}
