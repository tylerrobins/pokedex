package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreasList(pageURL *string) (LocationAreaListRes, error) {
    fullURL := baseURL + "/location-area" 
    if pageURL != nil {
        fullURL = *pageURL
    }
    req, err := http.NewRequest("GET", fullURL, nil) 
    if err != nil {
        return LocationAreaListRes{}, err
    }
    res, err := c.httpClient.Do(req) 
    if err != nil {
        return LocationAreaListRes{}, nil
    }   
    
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body) 
    if err != nil {
        return LocationAreaListRes{}, err
    }

    locationAreaList := LocationAreaListRes{}
    if err = json.Unmarshal(body, &locationAreaList); err != nil {
        return LocationAreaListRes{}, nil 
    }
    return locationAreaList, nil 
}
