package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (Pokemon, error) {
    fullURL := baseURL + "/pokemon/" + name    
    // check cache
    body, ok := c.cache.Get(fullURL)    
    if ok {
        fmt.Println("cache hit!")
        pokemon := Pokemon{}
        if err := json.Unmarshal(body, &pokemon); err != nil {
            return Pokemon{}, err
        }
    }
    fmt.Println("cache miss!")

    // request pokemon info
    req, err := http.NewRequest("GET", fullURL, nil) 
    if err != nil { return Pokemon{}, err }
    res, err := c.httpClient.Do(req) 
    if err != nil { return Pokemon{}, err }
    defer res.Body.Close()
    body, err = io.ReadAll(res.Body)
    if err != nil { return Pokemon{}, err }

    // update cache
    c.cache.Add(fullURL, body)

    // return pokemon info
    pokemon := Pokemon{}
    if err = json.Unmarshal(body, &pokemon); err != nil {
        return Pokemon{}, err 
    }
    return pokemon, nil
} 
