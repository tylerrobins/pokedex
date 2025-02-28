package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input:    "well Hello there  ",
            expected: []string{"well", "hello", "there"},
        },
        {
            input:    " how you DoiN",
            expected: []string{"how", "you", "doin"},
        },
        {
            input:    "cuz you had a bad day",
            expected: []string{"cuz", "you", "had", "a", "bad", "day"},
        },
        {
            input:    " hello its mE",
            expected: []string{"hello", "its", "me"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Actual length: %d and Expected length: %d are not equal", len(actual), len(c.expected))
            continue
        } 
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("Actual: %s and Expected: %s are not equal", word, expectedWord)
                continue
            } 
        }
    }
}
