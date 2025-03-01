package pokeapi

import (
	"testing"
	"time"
)

func TestGetFailing(t *testing.T) {
    pokeClient := NewClient(time.Hour)     
    fail := "fail"
    _, err := pokeClient.GetLocationAreasList(&fail)
    if err ==  nil {
        t.Error("Expected err")
    }
}
