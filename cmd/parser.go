package cmd

import (
	"fmt"
	"strings"
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
