package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    for {
        // Print prompt
        fmt.Fprint(os.Stdout, "$ ")

        // Wait for user input
        input, err := bufio.NewReader(os.Stdin).ReadString('\n')
        
        if err != nil {
            panic(err)
        }
        
        cmd := strings.TrimSpace(input)

        switch {
        case cmd == "exit 0": 
            os.Exit(0)
        case strings.HasPrefix(cmd, "echo"):
            args := strings.Split(cmd, " ")
            if args[0] == "echo" {
                fmt.Fprintln(os.Stdout, strings.Join(args[1:], " "))
            }
        default:
            fmt.Fprint(os.Stdout, cmd +": command not found\n")		
        }
    }
}