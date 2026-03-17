package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dwisarut/dealmaxxingCLI/internal/model"
)

func GetStoreData() []model.StoreLists {
	url := "https://www.cheapshark.com/api/1.0/stores"
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

	var stores []model.StoreLists
	err = json.Unmarshal(body, &stores)
	errorHandler(err)

	return stores
}

func GetGameFromId(id string) model.GetGameID {
	url := "https://www.cheapshark.com/api/1.0/games?id=" + id
	method := "GET"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	errorHandler(err)
	res, err := client.Do(req)

	if res.StatusCode == 404 || res.StatusCode == 403 {
		return model.GetGameID{}
	}

	errorHandler(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	errorHandler(err)
	var lists model.GetGameID
	err = json.Unmarshal(body, &lists)

	errorHandler(err)
	return lists
}

func GetGameIdentifier(name string) []model.GameIdentifier {
	url := "https://www.cheapshark.com/api/1.0/games?title=" + name
	method := "GET"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	errorHandler(err)
	res, err := client.Do(req)

	if res.StatusCode != 200 {
		return []model.GameIdentifier{}
	}

	errorHandler(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	errorHandler(err)

	var lists []model.GameIdentifier
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
