package main

import (
	"fmt"
)

func main() {
	todoList, err := LoadTodoList("todos.json")
	if err != nil {
		fmt.Println("Error Loading: ", err)
		return
	}

	ExecuteCLI(todoList)
}
