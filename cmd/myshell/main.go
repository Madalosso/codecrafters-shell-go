package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		commandName := input[:len(input)-1]
		fmt.Printf("%s: command not found\n", commandName)
	}
}
