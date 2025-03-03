package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
    catchThreshold := 50 
    pokemonName := args[0]
    if len(args) == 0 {
        fmt.Println("pokemon name is required to catch")
        return errors.New("No arg provided in catch command")
    } 
    fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName) 
    pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
    if err != nil {
        return err
    }
    catchChance := rand.IntN(pokemon.BaseExperience) 
    if catchChance > catchThreshold {
        fmt.Printf("failed to catch %s!\n", pokemonName)
        return nil 
    } 
    cfg.userPokedex[pokemonName] = pokemon
    fmt.Printf("%s was caught!!!\n", pokemon.Name)
    return nil
}
