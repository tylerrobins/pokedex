package main

import "fmt"

func commandMap(cfg *config) error {
    areasListRes, err := cfg.pokeapiClient.GetLocationAreasList(cfg.nextLocationAreaURL)
    if err != nil {
        return err
    }
    for _, area := range areasListRes.Results {
        fmt.Println(area.Name)
    }

    cfg.nextLocationAreaURL = &areasListRes.Next
    cfg.prevLocationAreaURL = &areasListRes.Previous
    return nil
}

func commandMapB(cfg *config) error {
    if cfg.prevLocationAreaURL == nil {
        fmt.Println("you're on the first page")
        return nil
    }
    areasListRes, err := cfg.pokeapiClient.GetLocationAreasList(cfg.prevLocationAreaURL)
    if err != nil {
        return err
    }
    for _, area := range areasListRes.Results {
        fmt.Println(area.Name)
    }
    
    cfg.nextLocationAreaURL = &areasListRes.Next
    if areasListRes.Previous != "" {
        cfg.prevLocationAreaURL = &areasListRes.Previous
    } else {
        cfg.prevLocationAreaURL = nil
    }
    return nil
}
