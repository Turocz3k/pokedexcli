package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startReplt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		clean_input := cleanInput(input)
		if len(clean_input) == 0 {
			continue
		}

		commandName := clean_input[0]

		switch commandsList[commandName].name {
		case "exit":
			err := commandsList[commandName].callback()
			if err != nil {
				fmt.Println(err)
			}
		case "help":
			err := commandsList[commandName].callback()
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
