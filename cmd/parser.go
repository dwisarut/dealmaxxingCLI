package cmd

import (
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/model"
)

func SearchParser(input string) model.SearchParseType {
	splitted := strings.Fields(input)
	var data model.SearchParseType

	switch {
	case len(splitted) > 2:
		title := strings.Join(splitted[1:], "%20")
		// title = strings.ToUpper(title)
		data = model.SearchParseType{Command: splitted[0], Arg: title}
		fmt.Printf("Data from long title game: %#v\n", data)
		return data

	case len(splitted) == 2:
		data = model.SearchParseType{Command: splitted[0], Arg: splitted[1]}
		fmt.Printf("Data: %#v\n", data)
		return data

	default:
		fmt.Println("Invalid command")
		return model.SearchParseType{}
	}

}

func QueryParser(input string) model.QueryParseType {
	splitted := strings.Fields(input)
	var data model.QueryParseType

	switch {
	case len(splitted) == 5:
		data = model.QueryParseType{
			Command:        splitted[0],
			SortBy:         splitted[1],
			Limit:          splitted[2],
			RatingSource:   splitted[3],
			MinRatingScore: splitted[4],
		}
		fmt.Printf("Default query data: %#v\n", data)
		return data

	case len(splitted) == 4:
		data = model.QueryParseType{
			Command:      splitted[0],
			SortBy:       splitted[1],
			Limit:        splitted[2],
			RatingSource: splitted[3],
		}
		fmt.Printf("Default query data: %#v\n", data)
		return data

	default:
		fmt.Println("Invalid command")
		return model.QueryParseType{}
	}
}
