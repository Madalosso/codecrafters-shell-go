package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commands = map[string]func(args []string){
	"exit": func(args []string) {
		codeStatus, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid exit code:", args[1], err)
		}
		os.Exit(codeStatus)
	},
	"echo": func(args []string) {
		fmt.Println(strings.Join(args[1:], " "))
	},
}

func main() {

	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")
		rawInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		inputLine := rawInput[:len(rawInput)-1]
		args := strings.Split(inputLine, " ")

		commandName := args[0]
		fn, ok := commands[commandName]
		if !ok {
			fmt.Printf("%s: command not found\n", commandName)
		} else {
			fn(args)
		}
	}
}
