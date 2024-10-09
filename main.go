package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	add    = "Add"
	show   = "Show"
	remove = "Remove"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage ./todo [command] [action] [due]")
		os.Exit(63)
	}

	command := args[1]

	switch command {
	case add:
		addTodo(args[2])
	case show:
		printMyTodos()
	case remove:

	default:
		fmt.Println("Invalid command use Add/Show/Remove")
		os.Exit(64)
	}

}

func addTodo(action string) {

	_, fileErr := os.Stat("./todo.store")

	if os.IsNotExist(fileErr) {
		file, err := os.Create("todo.store")

		if err != nil {
			fmt.Println("Could not create todo store")
			os.Exit(65)
		}

		writeAndClose(file, action)

		fmt.Println("Added new todo item")

	} else {
		file, err := os.OpenFile("./todo.store", os.O_APPEND, os.ModeAppend)
		if err != nil {
			fmt.Println("Cannot open todo store: ", err)
			os.Exit(66)
		}

		writeAndClose(file, action)

		fmt.Println("Added new todo item")
	}

}

func writeAndClose(file *os.File, action string) {
	defer file.Close()

	_, errWrite := file.WriteString("\n" + action)

	if errWrite != nil {
		fmt.Println("Error while adding item", errWrite.Error())
		os.Exit(67)
	}
}

func printMyTodos() {
	file, err := os.OpenFile("todo.store", os.O_RDONLY, os.ModePerm)

	if err != nil {
		fmt.Println("Cannot open todo store: ", err)
		os.Exit(66)
	}

	defer file.Close()

	buffer := make([]byte, 128)

	lines := ""

	for {
		_, readErr := file.Read(buffer)

		if readErr != nil {

			if readErr != io.EOF {
				fmt.Println("Error while reading file", readErr)
				os.Exit(67)
			}

			break
		}

		lines += string(buffer)
	}

	fmt.Printf("Your Todos\n\n")

	lns := strings.Split(lines, "\n")

	for i := 0; i < len(lns); i++ {
		if strings.Trim(lns[i], " ") != "" || lns[i] != "\n" {
			fmt.Printf("|%d %30s |\n", i+1, lns[i])
		}
	}
}
