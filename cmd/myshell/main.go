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
		case strings.HasPrefix(cmd, "type"):
			// if we enter type echo, it will print "echo is a shell command builtin"
			args := strings.Split(cmd, " ")
			if args[0] == "type" {
				fmt.Fprintln(os.Stdout, args[1] + " is a shell builtin")
			}
        default:
            fmt.Fprint(os.Stdout, cmd +": command not found\n")		
        }
    }
}