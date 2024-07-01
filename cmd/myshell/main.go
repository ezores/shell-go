package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"os/exec"
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
		/*case strings.HasPrefix(cmd, "type"):
			// if we enter type echo, it will print "echo is a shell command builtin"
			args := strings.Split(cmd, " ")
			if args[0] == "type" {
				if args[1] == "nonexistent" {
					fmt.Fprintln(os.Stdout, args[1] + ": not found")
				} else if args[1] == "nonexistentcommand" {
					fmt.Fprintln(os.Stdout, args[1] + ": not found")	
				} else if args[1] == "cat" {
					fmt.Fprintln(os.Stdout, args[1] + " is /bin/cat")
				} else {
					fmt.Fprintln(os.Stdout, args[1] + " is a shell builtin")
				}*/

		case strings.HasPrefix(cmd, "type"):
			args := strings.Split(cmd, " ")
			if args[0] == "type" {
				if args[1] == "echo" {
					fmt.Fprintln(os.Stdout, args[1] + " is a shell builtin")
				} else {
					path, err := exec.LookPath(args[1])
					if err != nil {
						fmt.Fprintln(os.Stdout, args[1] + ": not found")
					} else {
						fmt.Fprintln(os.Stdout, args[1] + " is " + path)
					}
				}
			}
        default:
            fmt.Fprint(os.Stdout, cmd +": command not found\n")		
        }
    }
}