package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type SupportedCommands string

const (
	Exit SupportedCommands = "exit"
	Echo SupportedCommands = "echo"
	Type SupportedCommands = "type"
)

func IsBuiltIn(commandName string) bool {
	switch SupportedCommands(commandName) {
	case Exit, Echo, Type:
		return true
	default:
		return false
	}
}

func exitHandler(args []string) {
	codeStatus, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid exit code:", args[1], err)
	}
	os.Exit(codeStatus)
}

func echoHandler(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func typeHandler(args []string) {
	cmdName := args[1]
	isBuiltIn := IsBuiltIn(cmdName)
	if isBuiltIn {
		fmt.Fprintf(os.Stdout, "%s is a shell %s\n", cmdName, "builtin")
	} else {

		// search in pathDirs for matching cmdName
		// refactor into own func
		path := os.Getenv("PATH")
		pathDirs := strings.Split(path, ":")
		pathToBin, err := checkOsCmd(pathDirs, cmdName)
		if err != nil {
			// fmt.Printf("%s: %s\n", cmdName, err)
			fmt.Fprintf(os.Stdout, "%s: not found\n", cmdName)

		} else {
			fmt.Fprintf(os.Stdout, "%s is %s\n", cmdName, pathToBin)
		}

	}
}

func checkOsCmd(pathDirs []string, cmd string) (string, error) {
	for _, dir := range pathDirs {
		filePath := filepath.Join(dir, cmd)
		_, err := os.Stat(filePath)
		if err == nil {
			return filePath, nil
		}
	}
	return "", fmt.Errorf("not found")

}

var commands = map[string]func(args []string){
	"exit": exitHandler,
	"echo": echoHandler,
	"type": typeHandler,
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
			path := os.Getenv("PATH")
			pathDirs := strings.Split(path, ":")
			pathToBin, err := checkOsCmd(pathDirs, commandName)
			if err != nil {
				// fmt.Printf("%s: %s\n", commandName, err)
				fmt.Fprintf(os.Stdout, "%s: not found\n", commandName)
			} else {
				cmd := exec.Command(pathToBin, args[1:]...)
				output, err := cmd.Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf(string(output))
			}
		} else {
			fn(args)
		}
	}
}
