package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreasList(pageURL *string) (LocationAreaListRes, error) {
    fullURL := baseURL + "/location-area?offset=0&limit=20"
 
    if pageURL != nil {
        fullURL = *pageURL
    }
    // cache logic
    body, ok := c.cache.Get(fullURL)
    if ok {
        fmt.Println("cache hit!")
        locationAreaList := LocationAreaListRes{}
        if err := json.Unmarshal(body, &locationAreaList); err != nil {
            return LocationAreaListRes{}, err 
        }
        return locationAreaList, nil
    } 
    fmt.Println("cache miss!")

    // request data
    req, err := http.NewRequest("GET", fullURL, nil) 
    if err != nil {
        return LocationAreaListRes{}, err
    }
    res, err := c.httpClient.Do(req) 
    if err != nil {
        return LocationAreaListRes{}, err 
    }   
    defer res.Body.Close()
    body, err = io.ReadAll(res.Body) 
    if err != nil {
        return LocationAreaListRes{}, err
    }
    
    // update cache
    c.cache.Add(fullURL, body)

    locationAreaList := LocationAreaListRes{}
    if err = json.Unmarshal(body, &locationAreaList); err != nil {
        return LocationAreaListRes{}, err 
    }
    return locationAreaList, nil 
}
