package pokeapi

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetLocationPokemon(locationName string) (pokemonEncountersResponse, error) {
	url := baseURL + "/location-area/" + locationName
	// Value is found in cache
	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := pokemonEncountersResponse{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return pokemonEncountersResponse{}, err
		}
		return pokemonResponse, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return pokemonEncountersResponse{}, err
	}
	if res.StatusCode != 200 {
		return pokemonEncountersResponse{}, fmt.Errorf("not a valid location")
	}
	defer res.Body.Close()
	pokemonResponse := pokemonEncountersResponse{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonEncountersResponse{}, err
	}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return pokemonEncountersResponse{}, err
	}
	c.cache.Add(url, data)
	return pokemonResponse, nil
}
