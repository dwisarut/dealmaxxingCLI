package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func SearchGameId(id string) {
	url := "https://www.cheapshark.com/api/1.0/deals?id=" + id
	method := "GET"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	errorHandler(err)
	res, err := client.Do(req)

	errorHandler(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	errorHandler(err)

	fmt.Println(string(body))
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// func getDealId(name string) {

// }
