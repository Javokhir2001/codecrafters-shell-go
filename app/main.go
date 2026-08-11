package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	br := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		readLine, err := br.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		command := strings.TrimSpace(readLine)
		if command == "exit" {
			os.Exit(0)
		}
		fmt.Println(command + ": command not found")
	}
}
