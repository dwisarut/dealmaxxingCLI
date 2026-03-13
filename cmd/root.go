package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
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

		case strings.HasPrefix(input, "search"):
			SearchHandler(reader, input)

		// case strings.HasPrefix(input, "get"):
		// 	var title string = CommonParser(input)
		// 	fmt.Println("Getting...")
		// 	api.GetDealFromTitle(title)

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
