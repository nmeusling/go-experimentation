package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl string) (locationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != "" {
		url = pageUrl
	}

	// Value is found in cache
	if val, ok := c.cache.Get(url); ok {
		locationsResponse := locationAreaResponse{}
		err := json.Unmarshal(val, &locationsResponse)
		if err != nil {
			return locationAreaResponse{}, err
		}
		return locationsResponse, nil

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
	c.cache.Add(url, data)
	return locationResponse, nil
}
