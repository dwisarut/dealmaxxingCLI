package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/cache"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/fatih/color"
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

	if val, ok := cachingID[cacheKey]; ok {
		fmt.Println(color.HiWhiteString("Identifier for"), color.HiCyanString(title))
		fmt.Println()
		fmt.Printf("%-81s %-10s\n", "Game title", "ID")

		for _, game := range val {
			fmt.Printf("%-90s %-10s\n", color.HiCyanString(game.Name), color.GreenString(game.GameID))
		}

		fmt.Println(strings.Repeat("_", 120))

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
		fmt.Printf("%-81s %-10s\n", "Game title", "ID")

		for _, list := range lists {
			fmt.Printf("%-90s %-10s\n", color.HiCyanString(list.Name), color.GreenString(list.GameID))
		}

		fmt.Println(strings.Repeat("_", 120))
		cache.IdentifierCaching(lists, cacheKey, cachingID)
	}
}

func titleValidation(data string) bool {
	if strings.Contains(data, "&") || strings.Contains(data, "=") || strings.Contains(data, "'") || strings.Contains(data, `"`) {
		return false
	}
	return true
}
