package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (LocationsResponse, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		var locations LocationsResponse
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return LocationsResponse{}, err
		}
		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationsResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResponse{}, err
	}

	c.cache.Add(url, data)

	var locations LocationsResponse
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsResponse{}, err
	}
	return locations, nil
}
