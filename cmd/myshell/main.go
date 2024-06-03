package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		rawInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		inputLine := rawInput[:len(rawInput)-1]
		args := strings.Split(inputLine, " ")

		commandName := args[0]
		switch commandName {
		case "exit":
			codeStatus, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid exit code:", args[1], err)
			}
			os.Exit(codeStatus)
		default:
			fmt.Printf("%s: command not found\n", commandName)
		}
	}
}
