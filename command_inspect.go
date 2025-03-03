package main

import (
    "errors"
    "fmt"
)

func commandInspect(cfg *config, args ...string) error {
    if len(args) == 0 {
        fmt.Println("pokemon name is required for inspect command")
        return errors.New("No arg give to inspect")
    }

    pokemonName := args[0]
    pokemon, ok := cfg.userPokedex[pokemonName] 
    if !ok {
        fmt.Println("you have not caught that pokemon")
        return nil
    }
    fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
`, pokemon.Name, pokemon.Height, pokemon.Weight)

    for _, stats := range pokemon.Stats {
        fmt.Printf(" -%s: %d\n", stats.Stat.Name, stats.BaseStat) 
    }
    fmt.Println("Types:")
    for _, types := range pokemon.Types {
        fmt.Printf(" - %s\n", types.Type.Name)
    }
    return nil 
}
