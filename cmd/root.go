package cmd

import (
	"flag"
	"fmt"
)

func InitCLI() {
	var search string
	var query string
	var command bool

	flag.StringVar(&search, "search", "", "Search game title")
	flag.StringVar(&query, "query", "", "Query deal")
	flag.BoolVar(&command, "cmd", false, "Show commands")

	initMessage()

	flag.Parse()

	if command {
		showCommand()
	}

	// if query != "" {
	// 	handleQuery(query)
	// }

	// if search != "" {
	// 	handleSearch(search)
	// }
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
