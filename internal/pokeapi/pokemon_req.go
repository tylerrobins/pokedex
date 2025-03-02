package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (PokemonDetailsRes, error) {
    fullURL := baseURL + "/pokemon/" + name    
    // check cache
    body, ok := c.cache.Get(fullURL)    
    if ok {
        fmt.Println("cache hit!")
        pokemonDetails := PokemonDetailsRes{}
        if err := json.Unmarshal(body, &pokemonDetails); err != nil {
            return PokemonDetailsRes{}, err
        }
    }
    fmt.Println("cache miss!")

    // request pokemon info
    req, err := http.NewRequest("GET", fullURL, nil) 
    if err != nil { return PokemonDetailsRes{}, err }
    res, err := c.httpClient.Do(req) 
    if err != nil { return PokemonDetailsRes{}, err }
    defer res.Body.Close()
    body, err = io.ReadAll(res.Body)
    if err != nil { return PokemonDetailsRes{}, err }

    // update cache
    c.cache.Add(fullURL, body)

    // return pokemon info
    pokemonDetails := PokemonDetailsRes{}
    if err = json.Unmarshal(body, &pokemonDetails); err != nil {
        return PokemonDetailsRes{}, err 
    }
    return pokemonDetails, nil
} 
