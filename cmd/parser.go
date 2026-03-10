package cmd

import (
	"fmt"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/model"
)

func CommonParser(input string) string {
	splitted := strings.Fields(input)
	var data string

	switch {
	case len(splitted) > 2:
		data = strings.Join(splitted[1:], "%20")
		return data

	case len(splitted) == 2:
		data = splitted[1]
		return data

	default:
		fmt.Println("Invalid command")
		return ""
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
