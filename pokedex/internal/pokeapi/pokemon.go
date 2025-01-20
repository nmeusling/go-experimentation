package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (pokemonResponse, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemon := pokemonResponse{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return pokemonResponse{}, err
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return pokemonResponse{}, err
	}
	defer res.Body.Close()
	pokemon := pokemonResponse{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonResponse{}, err
	}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemonResponse{}, err
	}
	c.cache.Add(url, data)
	return pokemon, nil
}
