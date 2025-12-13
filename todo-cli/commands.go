package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func ExecuteCLI(todoList *TodoList) {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// Main command
	cmd := os.Args[1]

	// Subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("title", "", "Task title")
	addPriority := addCmd.String("priority", "Low", "Task priority (Low/Medium/High)")
	addDue := addCmd.String("due", "", "Due date (YYYY-MM-DD)")

	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delID := delCmd.Int("id", 0, "ID of task to delete")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	completeID := completeCmd.Int("id", 0, "Mark task as complete")

	switch cmd {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTitle == "" || *addDue == "" {
			fmt.Println("Usage: todo.exe add -title=\"Buy Milk\" -priority=High -due=2025-12-13")
			return
		}

		priority := ParsePriority(*addPriority)
		dueDate, err := time.Parse("2006-01-02", *addDue)
		if err != nil {
			fmt.Println("Invalid date format. Use YYYY-MM-DD.")
			return
		}

		todoList.addTask(*addTitle, priority, dueDate)
		if err := todoList.saveData("todos.json"); err != nil {
			fmt.Println("Error saving: ", err)
		}

	case "delete":
		delCmd.Parse(os.Args[2:])

		if err := todoList.deleteTask(*delID); err != nil {
			fmt.Println(err)
		}

		if err := todoList.saveData("todos.json"); err != nil {
			fmt.Println("Error saving: ", err)
		}

	case "list":
		listCmd.Parse(os.Args[2:])
		todoList.printTodoList()

	case "complete":
		completeCmd.Parse(os.Args[2:])

		if err := todoList.markCompleted(*completeID); err != nil {
			fmt.Println(err)
		}

		if err := todoList.saveData("todos.json"); err != nil {
			fmt.Println("Error saving: ", err)
		}

	default:
		fmt.Println("Unknown command:", cmd)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Todo CLI Usage:")
	fmt.Println("  todo.exe add -title=\"Task\" -priority=High -due=2025-12-13")
	fmt.Println("  todo.exe delete -id=1")
	fmt.Println("  todo.exe complete -id=1")
	fmt.Println("  todo.exe list")
}
