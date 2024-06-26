package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	/** OLD VERSION **bufio.NewReader(os.Stdin).ReadString('\n') **/
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
}
