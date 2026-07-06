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

	errorHandler(err, 1)
	req.Header.Set("User-Agent", "dealmaxxingCLI/1.0")
	res, err := client.Do(req)

	errorHandler(err, 2)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	errorHandler(err, 3)

	var stores []model.StoreLists
	err = json.Unmarshal(body, &stores)
	errorHandler(err, 4)

	return stores
}

func GetGameFromId(id string) model.GetGameID {
	url := "https://www.cheapshark.com/api/1.0/games?id=" + id
	method := "GET"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	errorHandler(err, 1)
	req.Header.Set("User-Agent", "dealmaxxingCLI/1.0")
	res, err := client.Do(req)

	if res.StatusCode == 404 || res.StatusCode == 403 {
		return model.GetGameID{}
	}

	errorHandler(err, 2)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	errorHandler(err, 3)
	var lists model.GetGameID
	err = json.Unmarshal(body, &lists)

	errorHandler(err, 4)
	return lists
}

func GetGameIdentifier(name string) []model.GameIdentifier {
	url := "https://www.cheapshark.com/api/1.0/games?title=" + name
	method := "GET"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	errorHandler(err, 1)
	req.Header.Set("User-Agent", "dealmaxxingCLI/1.0")
	res, err := client.Do(req)

	if res.StatusCode != 200 {
		return []model.GameIdentifier{}
	}

	errorHandler(err, 2)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	errorHandler(err, 3)

	var lists []model.GameIdentifier
	err = json.Unmarshal(body, &lists)
	errorHandler(err, 4)

	return lists
}

func errorHandler(err error, num int) {
	if err != nil {
		fmt.Println("Line:", num)
		fmt.Println(err)
		panic(err)
	}
}
