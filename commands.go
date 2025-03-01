package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type ApiRes struct {
    Count    int    `json:"count"`
    Next     string `json:"next"`
    Previous string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

func getCommand() map[string]cliCommand {
    return map[string]cliCommand{
        "exit": {
            name:           "exit",
            description:    "Exit the Pokedex",
            callback:       commandExit,
        },
        "help": {
            name:           "help",
            description:    "Descibes how to use the Pokedex",
            callback:       commandHelp,
        },
        "map": {
            name:           "map",
            description:    "Gets the next 20 areas",
            callback:       commandMap,
        },
        "mapb": {
            name:           "mapb",
            description:    "Gets the previous 20 areas",
            callback:       commandMapB,
        },
        "explore": {
            name:           "explore",
            description:    "Gets pokemon in explored area",
            callback:       commandExplore,
        },
    }
}

func commandExit(cfg *config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config) error {
   fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
mapb: Gets list of previous 20 locations
`)
    return nil
}

func commandExplore(cfg *config) error {

}
