package main

import (
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
	"explore": {
		name:        "explore",
		description: "Show all pokemons in choosen location",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch choosen pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Check pokemon in your possesion",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Check pokemons in your possesion",
		callback:    commandPokedex,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func commandExit(c *config, s ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, s ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Show locations in Pokemon world forward")
	fmt.Println("mapb: Show locations in Pokemon world backward")
	fmt.Println("explore: Show all pokemons in choosen location")
	fmt.Println("catch: Catch choosen pokemon")
	fmt.Println("inspect: Check pokemon in your possesion")
	fmt.Println("pokedex: Check pokemons in your possesion")
	fmt.Println("exit: Exit the Pokedex")

	/*
		for _, val := range commandsList {
			fmt.Println(val.description)
		}
	*/
	return nil
}

func commandMap(c *config, s ...string) error {
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

func commandMapb(c *config, s ...string) error {
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

func commandExplore(c *config, place ...string) error {
	resp, err := c.pokeapiClient.GetLocation(place[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + resp.Name + "...")
	fmt.Println("Found Pokemon:")

	for _, v := range resp.PokemonEncounters {
		fmt.Println("- " + v.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *config, name ...string) error {
	fmt.Println("Throwing a Pokeball at " + name[0] + "...")

	resp, err := c.pokeapiClient.GetPokemon(name[0])
	if err != nil {
		return err
	}

	chance := rand.Intn(100)

	if (resp.BaseExperience / 10) <= chance {
		c.pokedex[name[0]] = resp
		fmt.Println(name[0] + " was caught!")
		return nil
	}

	fmt.Println(name[0] + " escaped!")
	return nil
}

func commandInspect(c *config, name ...string) error {
	v, ok := c.pokedex[name[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name: " + v.Name)
	fmt.Printf("Height: %d\n", v.Height)
	fmt.Printf("Weight: %d\n", v.Weight)

	fmt.Println("Stats:")
	for _, v := range v.Stats {
		fmt.Printf(" -%s: %d\n", v.Stat.Name, v.BaseStat)
	}

	fmt.Println("Types:")
	for _, v := range v.Types {
		fmt.Println(" -" + v.Type.Name)
	}

	return nil
}

func commandPokedex(c *config, s ...string) error {
	if len(c.pokedex) < 1 {
		fmt.Println("Pokdex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, v := range c.pokedex {
		fmt.Println(" - " + v.Name)
	}

	return nil
}
