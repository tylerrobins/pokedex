package main

import (
	"fmt"
	"os"
	"sort"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
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
        "catch": {
            name:           "catch",
            description:    "Attemptes to catch pokemon",
            callback:       commandCatch,
        },
        "inspect": {
            name:           "inspect",
            description:    "Inspects the caught pokemon",
            callback:       commandInspect,
        },
        "pokedex": {
            name:           "pokedex",
            description:    "Lists all pokemon in your pokedex",
            callback:       commandPokedex,
        },
    }
}

func commandExit(cfg *config, args ...string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config, args ...string) error {
    commands := getCommand()

    // get command list
    var commandList []string 
    for name := range commands {
        commandList = append(commandList, name)
    }
    sort.Strings(commandList)

    fmt.Print("Welcome to the Pokedex!\nUsage:\n")
    
    // find longest name for formatting
    maxLen := 0
    for _, name := range commandList {
        if len(name) > maxLen { maxLen = len(name) }
    }

    // print commands 
    for _, name := range commandList {
        com := commands[name]
        fmt.Printf("- %-*s      %s\n", maxLen, com.name, com.description) 
    } 
    fmt.Println("")
    return nil
}
