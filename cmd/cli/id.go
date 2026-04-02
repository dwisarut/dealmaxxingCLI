package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/cache"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func IDHandler(reader *bufio.Reader, input string, cachingID map[string][]model.GameIdentifier) {
	title := CommonParser(input)
	isValid := titleValidation(title)

	if !isValid {
		color.HiRed("Invalid input! Try typing a proper name again :D")
		return
	}

	fmt.Println()
	color.Green("Getting ID...")
	fmt.Println()

	cacheKey := cache.MakeCachingKey(title)

	headerFmt := color.New(color.FgHiWhite, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgHiCyan).SprintfFunc()

	tbl := table.New("Game title", "ID")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	if val, ok := cachingID[cacheKey]; ok {

		for _, game := range val {
			tbl.AddRow(game.Name, game.GameID)
		}

		tbl.Print()
		fmt.Println(strings.Repeat("-", 10))

	} else {
		lists := api.GetGameIdentifier(title)

		if len(lists) == 0 {
			fmt.Println(color.HiRedString("No games founded."))
			fmt.Println("Exiting...")
			fmt.Println()
			return
		}

		fmt.Println(color.HiWhiteString("Identifier for"), color.HiCyanString(title))
		fmt.Println()

		for _, list := range lists {
			tbl.AddRow(list.Name, list.GameID)
		}

		tbl.Print()
		fmt.Println(strings.Repeat("-", 10))
		cache.IdentifierCaching(lists, cacheKey, cachingID)
	}
}

func titleValidation(data string) bool {
	if strings.Contains(data, "&") || strings.Contains(data, "=") || strings.Contains(data, "'") || strings.Contains(data, `"`) {
		return false
	}
	return true
}
