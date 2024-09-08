package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cliUi()
}

func cliUi() {
	params := os.Args

	if params != nil && params[1] == "add" && params[2] != "" {
		if len(params) != 3 {
			fmt.Println("Usage: go run main.go <add> <\"description\">")
			return
		}

		AddNewTask(params[2])
	}

	if params != nil && params[1] == "delete" {
		if len(params) != 3 {
			fmt.Println("Usage: go run main.go <delete> <integerID>")
			return
		}

		intArg, err := strconv.Atoi(params[2])
		if err != nil {
			fmt.Println("The second argument must be an integer")
		}

		DeleteTask(intArg)
		fmt.Printf("Task %d deleted", intArg)
	}

	if params != nil && params[1] == "update" && params[3] != "" {
		if len(params) != 4 {
			fmt.Println("Usage: go run main.go <update> <integerID> <\"newDescription\">")
			return
		}
		intArg, err := strconv.Atoi(params[2])
		if err != nil {
			fmt.Println("The second argument must be an integer")
		}

		newDescription := os.Args[3]

		UpdateTask(intArg, newDescription)
	}

	// marking a task an in-progress, done or todo
	if params != nil && strings.Contains(params[1], "mark") {
		if len(params) != 3 {
			fmt.Println("Usage: go run main.go <update> <integerID> <\"newDescription\">")
			return
		}
		intArg, err := strconv.Atoi(params[2])
		if err != nil {
			fmt.Println("The second argument must be an integer")
		}

		switch {
		case strings.Contains(params[1], "todo"):
			UpdateStatus(TODO, intArg)
		case strings.Contains(params[1], "in-progress"):
			UpdateStatus(INPROGRESS, intArg)
		case strings.Contains(params[1], "done"):
			UpdateStatus(DONE, intArg)
		}
	}

	if params != nil && params[1] == "list" {
		if len(params) == 2 {
			file, err := ReadFile()
			if err != nil {
				return
			}
			fmt.Println(file)
		}
		if len(params) == 3 {
			fmt.Println(ListTasksByStatus(Status(params[2])))
		}
	}
}
