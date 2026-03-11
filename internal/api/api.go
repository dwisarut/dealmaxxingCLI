package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dwisarut/dealmaxxingCLI/internal/model"
)

func GetGameFromId(id string) {
	url := "https://www.cheapshark.com/api/1.0/games?id=" + id
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

func GetDealFromTitle(name string) []model.SearchGameID {
	url := "https://www.cheapshark.com/api/1.0/games?title=" + name
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

	var lists []model.SearchGameID
	err = json.Unmarshal(body, &lists)
	errorHandler(err)

	return lists
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
