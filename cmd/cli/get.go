package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/cache"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/dwisarut/dealmaxxingCLI/internal/service"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func GetHandler(reader *bufio.Reader, input string, storeIndex map[string]string, cachingGet map[string]model.GetGameID) {
	id := CommonParser(input)
	isValid := idValidation(id)

	if !isValid {
		color.HiRed("Invalid input! Try typing a proper name again :D")
		return
	}

	config := api.ReadEnv()
	client := api.NewService(config)

	kvData := api.GetKV(config, client)
	kvCount := kvData.Count
	fmt.Println("Current KV count: ", kvCount)

	fmt.Println()
	color.Green("Getting...")
	fmt.Println()

	cacheKey := cache.MakeCachingKey(id)

	headerFmt := color.New(color.FgHiWhite, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgHiBlue).SprintfFunc()

	tbl := table.New("Store", "Price ($)", "Link")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	if val, ok := cachingGet[cacheKey]; ok {
		fmt.Println(color.HiCyanString(val.Info.Title))
		fmt.Println()

		for _, list := range val.Deals {
			tbl.AddRow(list.StoreName, list.Price, list.Redirect)
		}

		tbl.Print()
		fmt.Println(strings.Repeat("-", 10))

	} else {
		var game model.GetGameID = api.GetGameFromId(id)

		if len(game.Deals) == 0 {
			fmt.Println(color.HiRedString("No deal currently available."))
			fmt.Println("Exiting...")
			fmt.Println()
			return
		}

		var displayList model.GetGameID = service.MakeGetRedirect(game)

		// TODO: MAKE URL SHORTENER HERE
		// Function that take kvCount and encode the URL

		displayList = service.MatchGetStore(displayList, storeIndex)

		fmt.Println(color.HiCyanString(displayList.Info.Title))
		fmt.Println()

		for _, list := range displayList.Deals {
			tbl.AddRow(list.StoreName, list.Price, list.Redirect)
		}

		tbl.Print()
		fmt.Println(strings.Repeat("-", 10))

		// Cache the shortened URL around here!
		cache.GetCaching(displayList, cacheKey, cachingGet)
	}
}

func idValidation(data string) bool {
	if strings.Contains(data, "&") || strings.Contains(data, "=") || strings.Contains(data, "'") || strings.Contains(data, `"`) {
		return false
	}
	return true
}
