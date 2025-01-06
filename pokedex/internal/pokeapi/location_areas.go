package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas(pageUrl string) (locationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != "" {
		url = pageUrl
	}
	res, err := http.Get(url)
	if err != nil {
		return locationAreaResponse{}, err
	}
	defer res.Body.Close()
	locationResponse := locationAreaResponse{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaResponse{}, err
	}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return locationAreaResponse{}, err
	}
	for _, l := range locationResponse.Results {
		fmt.Println(l.Name)
	}
	return locationResponse, nil
}
