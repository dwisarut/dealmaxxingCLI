package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func searchGameId(id string) {
	url := "https://www.cheapshark.com/api/1.0/deals?id=" + id
	method := "GET"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	ErrorHandler(err)
	res, err := client.Do(req)

	ErrorHandler(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	ErrorHandler(err)

	fmt.Println(string(body))
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// func getDealId(name string) {

// }
