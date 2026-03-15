package cmd

import (
	"net/url"
	"strings"

	"github.com/fatih/color"
)

func CommonParser(input string) string {
	splitted := strings.Fields(input)
	var data string

	switch {
	case len(splitted) > 2:
		data = strings.Join(splitted[1:], "")
		data = url.QueryEscape(data)
		return data

	case len(splitted) == 2:
		data = splitted[1]
		return data

	case len(splitted) < 2:
		color.HiRed("Invalid parameters, please search again!")
		return ""

	default:
		color.HiRed("Invalid command")
		return ""
	}

}
