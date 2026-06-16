package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Location struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocation(pageURL string) (Location, error) {
	const baseURL = "https://pokeapi.co/api/v2/location-area/"
	url := baseURL + pageURL

	if val, ok := c.cache.Get(url); ok {
		var pokemons Location
		err := json.Unmarshal(val, &pokemons)
		if err != nil {
			return Location{}, err
		}
		return pokemons, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)
	var pokemons Location
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return Location{}, err
	}

	return pokemons, nil
}
