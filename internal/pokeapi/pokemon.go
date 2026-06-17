package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Name           string `json:"name"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	const baseURL = "https://pokeapi.co/api/v2/pokemon/"
	url := baseURL + name

	if val, ok := c.cache.Get(url); ok {
		var pokemon Pokemon
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
