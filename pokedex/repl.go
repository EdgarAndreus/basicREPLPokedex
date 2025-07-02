package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/EdgarAndreus/pokedexcli/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 locations areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Displays the previous 20 locations areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "List the pokemon located in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a pokemon and adds it to the Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in the Pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	var config Config
	config.cache = pokecache.NewCache(5 * time.Second)
	config.Next = ""
	config.Previous = ""
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		cInput := cleanInput(reader.Text())
		if len(cInput) == 0 {
			continue
		}
		command, exist := getCommands()[cInput[0]]
		if exist {
			err := command.callback(&config, cInput[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
