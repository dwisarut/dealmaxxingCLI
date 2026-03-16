package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/api"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/dwisarut/dealmaxxingCLI/internal/service"
	"github.com/fatih/color"
)

func GetHandler(reader *bufio.Reader, input string, storeIndex map[string]string) {
	var id string = CommonParser(input)

	if id == "" {
		return
	}

	fmt.Println()
	color.Green("Getting...")
	fmt.Println()

	var game model.GetGameID = api.GetGameFromId(id)

	if len(game.Deals) == 0 {
		fmt.Println(color.HiRedString("No deal currently available."))
		fmt.Println("Exiting...")
		fmt.Println()
		return
	}

	var displayList model.GetGameID = service.MakeGetRedirect(game)
	displayList = service.MatchGetStore(displayList, storeIndex)

	color.HiWhite("Currently available deals:")
	fmt.Println()
	for _, list := range displayList.Deals {
		fmt.Println(color.HiCyanString(displayList.Info.Title))
		fmt.Println(color.HiBlueString("Store:"), color.HiBlueString(list.StoreName))
		fmt.Println(color.HiYellowString("Prices:"), color.YellowString(list.Price), color.YellowString("$"))
		fmt.Println(color.HiWhiteString("Link:"), color.GreenString(list.Redirect))
		fmt.Println()
	}

	fmt.Println(strings.Repeat("_", 120))
}
