package cli

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
		isValid := inputValidation(data)

		if !isValid {
			return ""
		}

		return data

	case len(splitted) == 2:
		data = splitted[1]
		isValid := inputValidation(data)

		if !isValid {
			return ""
		}

		return data

	case len(splitted) < 2:
		color.HiRed("Invalid parameters, please search again!")
		return ""

	default:
		color.HiRed("Invalid command")
		return ""
	}

}

func inputValidation(data string) bool {
	if strings.Contains(data, "&") || strings.Contains(data, "=") || strings.Contains(data, "'") || strings.Contains(data, `"`) {
		return false
	}
	return true
}
