package main

import (
	"fmt"
	"os"
)

const (
	add    = "Add"
	show   = "Show"
	remove = "Remove"
)

func main() {

	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage ./todo [command] [action] [due]")
		os.Exit(63)
	}

	command := args[1]

	switch command {
	case add:
		addTodo(args[2])
	case show:

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

		file.WriteString(action)

		defer file.Close()
	} else {
		file, err := os.Open("./todo.store")
		if err != nil {
			fmt.Println("Cannot open todo store: ", err)
			os.Exit(66)
		}

		file.WriteString(action)

		defer file.Close()
	}

}