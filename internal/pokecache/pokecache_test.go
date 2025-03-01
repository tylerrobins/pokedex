package pokecache

import (
	"testing"
	"time"
    "fmt"
)

func TestCreateCache(t *testing.T) {
    c :=  NewCache(time.Second)
    if c.cache == nil {
        t.Error("Cache wasn't create successfully")
    }
}

func TestAddGetCache(t *testing.T) {
    cache := NewCache(time.Millisecond)
    cases := []struct {
        key string
        val []byte
        }{
        {
			key: "Key1",
			val: []byte("val1"),
		},
		{
			key: "KeY2",
			val: []byte("vAl2"),
		},
		{
			key: "",
			val: []byte("VAL3"),
		},
    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
            cache.Add(c.key, c.val)    
            val, ok := cache.Get(c.key)
            if !ok {
                t.Error("Expected to find key")
                return
            }
            if string(val) != string (c.val) {
                t.Error("Expected for find correct value")
                return
            }
        })
    }
}

func TestReapCahce(t *testing.T) {
    cache := NewCache(time.Millisecond)
    cases := []struct {
        key string
        val []byte
        }{
        {
			key: "Key1",
			val: []byte("val1"),
		},
		{
			key: "KeY2",
			val: []byte("vAl2"),
		},
		{
			key: "",
			val: []byte("VAL3"),
		},
    }
    
    for i, c := range cases {
        t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
            cache.Add(c.key, c.val)    
            val, ok := cache.Get(c.key)
            if !ok {
                t.Error("Expected to find key")
                return
            }
            if string(val) != string (c.val) {
                t.Error("Expected for find correct value")
                return
            }
        })
    }

    time.Sleep(time.Millisecond * 10)
    if len(cache.cache) != 0 {
        t.Error("Expected cache to be reaped")
        return
    }
}
