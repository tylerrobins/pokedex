package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
    if len(args) == 0 {
        fmt.Println("location `name` or `id` is required with explore command")
        return errors.New("No args provided")
    }
	fmt.Printf("Exploring %s\n", args[0])
    locationAreaInfoRes, err := cfg.pokeapiClient.GetLocationInfo(args[0])
    if err != nil {
        return err
    }
    fmt.Println("Found Pokemon:")
    for _, area := range locationAreaInfoRes.PokemonEncounters {
       fmt.Printf(" - %s\n",area.Pokemon.Name) 
    }
	return nil
}
