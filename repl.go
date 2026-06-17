package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleanInput := cleanInput(input)
		if len(cleanInput) == 0 {
			continue
		}

		commandName := cleanInput[0]

		switch commandsList[commandName].name {
		case "exit":
			err := commandsList[commandName].callback(c, "")
			if err != nil {
				fmt.Println(err)
			}
		case "help":
			err := commandsList[commandName].callback(c, "")
			if err != nil {
				fmt.Println(err)
			}
		case "map":
			err := commandsList[commandName].callback(c, "")
			if err != nil {
				fmt.Println(err)
			}
		case "mapb":
			err := commandsList[commandName].callback(c, "")
			if err != nil {
				fmt.Println(err)
			}
		case "explore":
			err := commandsList[commandName].callback(c, cleanInput[1])
			if err != nil {
				fmt.Println(err)
			}
		case "catch":
			err := commandsList[commandName].callback(c, cleanInput[1])
			if err != nil {
				fmt.Println(err)
			}
		case "inspect":
			err := commandsList[commandName].callback(c, cleanInput[1])
			if err != nil {
				fmt.Println(err)
			}
		case "pokedex":
			err := commandsList[commandName].callback(c, "")
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	result := strings.Fields(output)
	return result
}
