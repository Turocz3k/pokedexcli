package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commandsList = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Show locations in Pokemon world forward",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Show locations in Pokemon world backward",
		callback:    commandMapb,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Show locations in Pokemon world")
	fmt.Println("exit: Exit the Pokedex")

	/*
		for _, val := range commandsList {
			fmt.Println(val.description)
		}
	*/
	return nil
}

///////////////////////Bellow everything to package internal/pokeapi  //////////////////////

type locationResApi struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Next != nil {
		url = *c.Next
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Http.Get() doesnt work!")
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("%v", res.StatusCode)
	}

	var locations locationResApi
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&locations)
	if err != nil {
		return fmt.Errorf("Decode doesnt work!")
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for i := range 20 {
		fmt.Printf("%v\n", locations.Results[i].Name)
	}
	return nil
}

func commandMapb(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Previous != nil {
		url = *c.Previous
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Http.Get() doesnt work!")
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("%v", res.StatusCode)
	}

	var locations locationResApi
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&locations)
	if err != nil {
		return fmt.Errorf("Decode doesnt work!")
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for i := range 20 {
		fmt.Printf("%v\n", locations.Results[i].Name)
	}

	return nil
}
