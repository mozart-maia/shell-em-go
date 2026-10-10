package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			return
		}

		inp_list := strings.Split(input, " ")
		command := strings.TrimSpace(inp_list[0])

		if command == "exit" {
			os.Exit(0)
		}
		if command == "echo" {
			inp_list = inp_list[1:]
			echo_output := strings.Join(inp_list, " ")
			fmt.Print(echo_output)

		} else {
			fmt.Println(command + ": command not found")
		}

	}

}
