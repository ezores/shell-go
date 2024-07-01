package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")
	for{
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		
		if(err != nil) {
			panic(err)
		}
		
		cmd := strings.TrimSpace(input)

		switch cmd {
		case "exit 0": 
			os.Exit(0)
		case "echo"
				args := strings.Split(cmd, " ")
				if args[0] == "echo" {
					fmt.Fprintln(os.Stdout, strings.Join(args[1:], " "))
				}
		default:
			fmt.Fprint(os.Stdout, cmd +": command not found\n")		
	}
	
}
}
