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

	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')

		if err != nil {
			return
		}

		comlist := strings.Split(command, " ")
		output := strings.TrimSpace(comlist[0])

		if output == "exit" {
			os.Exit(0)
		}
		fmt.Println(output + ": command not found")
	}

}
