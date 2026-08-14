package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func executableExists(path string, filename string) (bool, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	for _, file := range files {
		fname := file.Name()
		if fname == filename {

			finfo, err := os.Lstat(path + "/" + fname)
			if err != nil {
				return false, err
			}
			isExecutable := finfo.Mode().Perm()&0111 != 0
			return isExecutable, nil
		}

	}

	return false, nil
}
func main() {
	// TODO: Uncomment the code below to pass the first stage
	pathVar := os.Getenv("PATH")
	pathDirs := strings.Split(pathVar, ":")
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
				found := false
				for _, path := range pathDirs {
					isExecutable, err := executableExists(path, cmd)
					if err != nil {
						fmt.Println(err)
						break
					}
					if isExecutable {
						fmt.Println(cmd + " is " + path + "/" + cmd)
						found = true
						break
					}
				}

				if !found {
					fmt.Println(cmd + ": not found")
				}
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
