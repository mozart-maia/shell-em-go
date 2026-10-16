package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	built_ins := []string{"echo", "type", "exit"}

	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			return
		}

		input = strings.TrimSpace(input)
		inp_list := strings.Fields(input)
		command := strings.TrimSpace(inp_list[0])

		if command == "exit" {
			os.Exit(0)
		}

		if command == "echo" {
			inp_list = inp_list[1:]
			echo_output := strings.Join(inp_list, " ")
			fmt.Println(echo_output)
			continue
		}

		if command == "type" {
			type_com(inp_list, built_ins)
			continue
		}

		args := []string(inp_list[1:])
		cmd := exec.Command(command, args...)
		resultado, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(command + ": command not found")
			continue
		}

		fmt.Print(string(resultado))

	}
}

func built_ins_contain(b_ins []string, arg string) bool {
	for i := range b_ins {
		if b_ins[i] == arg {
			return true
		}
	}
	return false
}

func type_com(inp_list []string, built_ins []string) {
	com_described := inp_list[1]
	is_com := built_ins_contain(built_ins, com_described)
	if is_com == true {
		fmt.Println(com_described + " is a shell builtin")
	} else {

		path, err := exec.LookPath(com_described)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(com_described + ": not found")
		} else {
			fmt.Println(com_described + " is " + path)
		}
	}
}
