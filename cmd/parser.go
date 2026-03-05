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
		title := strings.Join(splitted[1:], " ")
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
