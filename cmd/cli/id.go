package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/fatih/color"
)

func IDHandler(reader *bufio.Reader, input string) {
	var title string = CommonParser(input)

	if title == "" {
		return
	}

	fmt.Println()
	color.Green("Getting ID...")
	fmt.Println()

	var lists []model.GameIdentifier = api.GetGameIdentifier(title)

	if len(lists) == 0 || len(lists) == 60 {
		fmt.Println(color.HiRedString("No games founded."))
		fmt.Println("Exiting...")
		fmt.Println()
		return
	}

	fmt.Println(color.HiWhiteString("Identifier for"), color.HiCyanString(title))
	fmt.Println()
	fmt.Printf("%-81s %-10s\n", "Game title", "ID")

	for _, list := range lists {
		fmt.Printf("%-90s %-10s\n", color.HiCyanString(list.Name), color.GreenString(list.GameID))
	}

	fmt.Println(strings.Repeat("_", 120))
}
