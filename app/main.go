package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	builtinCommands := map[string]bool{
		"echo": true,
		"exit": true,
		"type": true,
	}
	br := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		readLine, err := br.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		commandLine := strings.TrimSpace(readLine)
		commandLineSplit := strings.Split(commandLine, " ")
		command := commandLineSplit[0]
		switch command {
		case "echo":
			fmt.Println(commandLine[5:])
		case "exit":
			os.Exit(0)
		case "type":
			cmd := commandLineSplit[1]
			_, ok := builtinCommands[cmd]
			if ok {
				fmt.Println(cmd + " is a shell builtin")
			} else {
				fmt.Println(cmd + ": not found")
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
