package main

import (
    "strings"
    "fmt"
    "bufio"
    "os"
)


func startRepl(cfg *config){
    avalCommands := getCommand()
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ") 
        scanner.Scan()
        input := scanner.Text()
        cleaned := cleanInput(input)
        if len(cleaned) == 0 {
            continue
        }
        
        command, ok := avalCommands[cleaned[0]]
        if !ok {
            fmt.Printf("Invalid command: %s, try `help` for a list of commands\n", cleaned[0])
            continue
        }
        command.callback(cfg, cleaned[1:]...)
    }
}

func cleanInput(test string) []string {
    strLower := strings.ToLower(test)
    words := strings.Fields(strLower)
    return words 
 }
