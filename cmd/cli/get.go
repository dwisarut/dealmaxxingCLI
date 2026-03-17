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
)

func GetHandler(reader *bufio.Reader, input string, storeIndex map[string]string, cachingGet map[string]model.GetGameID) {
	id := CommonParser(input)
	isValid := idValidation(id)

	if !isValid {
		color.HiRed("Invalid input! Try typing a proper name again :D")
		return
	}

	fmt.Println()
	color.Green("Getting...")
	fmt.Println()

	cacheKey := cache.MakeCachingKey(id)

	if val, ok := cachingGet[cacheKey]; ok {
		color.HiYellow("Cache Hit")
		color.HiWhite("Currently available deals:")
		fmt.Println()
		fmt.Println(color.HiCyanString(val.Info.Title))

		for _, list := range val.Deals {
			fmt.Println(color.HiBlueString("Store:"), color.HiBlueString(list.StoreName))
			fmt.Println(color.HiYellowString("Prices:"), color.YellowString(list.Price), color.YellowString("$"))
			fmt.Println(color.HiWhiteString("Link:"), color.GreenString(list.Redirect))
			fmt.Println()
		}

	} else {
		var game model.GetGameID = api.GetGameFromId(id)

		if len(game.Deals) == 0 {
			fmt.Println(color.HiRedString("No deal currently available."))
			fmt.Println("Exiting...")
			fmt.Println()
			return
		}

		var displayList model.GetGameID = service.MakeGetRedirect(game)
		displayList = service.MatchGetStore(displayList, storeIndex)

		color.HiYellow("API Hit")
		color.HiWhite("Currently available deals:")
		fmt.Println()
		fmt.Println(color.HiCyanString(displayList.Info.Title))

		for _, list := range displayList.Deals {
			fmt.Println(color.HiBlueString("Store:"), color.HiBlueString(list.StoreName))
			fmt.Println(color.HiYellowString("Prices:"), color.YellowString(list.Price), color.YellowString("$"))
			fmt.Println(color.HiWhiteString("Link:"), color.GreenString(list.Redirect))
			fmt.Println()
		}

		fmt.Println(strings.Repeat("_", 120))
		cache.GetCaching(displayList, cacheKey, cachingGet)
	}
}

func idValidation(data string) bool {
	if strings.Contains(data, "&") || strings.Contains(data, "=") || strings.Contains(data, "'") || strings.Contains(data, `"`) {
		return false
	}
	return true
}
