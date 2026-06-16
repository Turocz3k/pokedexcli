package main

import (
	"fmt"
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
	fmt.Println("map: Show locations in Pokemon world forward")
	fmt.Println("mapb: Show locations in Pokemon world backward")
	fmt.Println("exit: Exit the Pokedex")

	/*
		for _, val := range commandsList {
			fmt.Println(val.description)
		}
	*/
	return nil
}

func commandMap(c *config) error {
	resp, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}
	c.nextLocationsURL = resp.Next
	c.prevLocationsURL = resp.Previous
	for _, v := range resp.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	resp, err := c.pokeapiClient.ListLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}
	c.nextLocationsURL = resp.Next
	c.prevLocationsURL = resp.Previous
	for _, v := range resp.Results {
		fmt.Println(v.Name)
	}
	return nil
}
