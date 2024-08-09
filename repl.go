package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/misterlister/pokedex_cli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	locationPage     int
}

func startRepl() {
	pokeClient := pokeapi.NewClient(10*time.Second, 30*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		locationPage:  0,
	}
	validCommands := getCommands()
	err := commandHelp(cfg, "")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		arg := ""
		if len(words) > 1 {
			arg = words[1]
		}
		fmt.Println()
		command, exists := validCommands[commandName]
		if exists {
			err := command.callback(cfg, arg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Invalid command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays valid commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex application",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next page of map locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all the pokemon found in a specified area (eg. 'explore mt-coronet-2f')",
			callback:    commandExplore,
		},
	}
}
