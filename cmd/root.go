package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/service"
	"github.com/fatih/color"
)

func InitCLI() {
	reader := bufio.NewReader(os.Stdin)

	initMessage()
	storesMap := fetchStoreMap()

	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		case strings.HasPrefix(input, "cmd"):
			showCommand()

		case strings.HasPrefix(input, "search"):
			SearchHandler(reader, input, storesMap)

		case strings.HasPrefix(input, "get"):
			GetHandler(reader, input, storesMap)

		case input == "exit":
			fmt.Println("Exiting DealmaxxingCLI...")
			return

		default:
			fmt.Println("Unknown command")
		}
	}

}

func showCommand() {
	fmt.Println()
	fmt.Println("Here's an available commands:")
	fmt.Printf("%s\t\t%s\n", color.GreenString("search"), "Searching for a specific game")
	fmt.Printf("%s\t\t%s\n", color.GreenString("get"), "Find a specific game with ID for all available deals")
}

func initMessage() {
	fmt.Println()
	fmt.Println(color.HiGreenString("DealmaxxingCLI"))
	fmt.Println("Your trusty tools for finding cheap game!")
	fmt.Println("Type", color.HiYellowString("cmd"), "to find an available commands.")
}

func fetchStoreMap() map[string]string {
	stores := api.GetStoreData()
	storesMap := service.IndexingStore(stores)
	return storesMap
}
