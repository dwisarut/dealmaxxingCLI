package cmd

import (
	"flag"
	"fmt"

	"github.com/dwisarut/dealmaxxingCLI/internal/model"
)

func InitCLI() {
	var inputData model.InputDataCLI

	flag.StringVar(&inputData.Search, "search", "", "Search game title")
	flag.StringVar(&inputData.Query, "query", "", "Query deal")
	flag.BoolVar(&inputData.Command, "cmd", false, "Show commands")
	flag.IntVar(&inputData.Top, "top", 10, "Top N deals currently")
	flag.IntVar(&inputData.MaxPrice, "maxprice", 15, "Maximum price filter")

	initMessage()

	flag.Parse()

	switch {
	case inputData.Command:
		showCommand()

	// case inputData.Query != "":
	// 	handleQuery(inputData.Query)

	// case inputData.Query != "":
	// 	handleSearch(inputData.Search)

	default:
		fmt.Println("No command provided")
	}

}

func showCommand() {
	fmt.Println("Here's an available commands:")
	fmt.Println("-search					Searching for a specific game")
	fmt.Println("-query					Query lists of game with specific parameter")
}

func initMessage() {
	fmt.Println("DealmaxxingCLI")
	fmt.Println("Insert your command here.")
	fmt.Println("Type -cmd to find an available commands!")
}
