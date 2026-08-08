package main

import (
	"fmt"
	"os"
	"bufio"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")
	br := bufio.NewReader(os.Stdin)
	for {
		readLine := br.ReadString('\n')
		fmt.Printf("%s: command not found", readLine)
	}
}
