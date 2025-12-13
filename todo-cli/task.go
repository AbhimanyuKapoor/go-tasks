package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type Priority int
type TodoList []Task

const (
	Low Priority = iota
	Medium
	High
)

type Task struct {
	Title     string
	Priority  Priority
	DueDate   time.Time
	Completed bool
}

func (todoList *TodoList) addTask(title string, priority Priority, dueDate time.Time) {
	task := Task{
		Title:     title,
		Priority:  priority,
		DueDate:   dueDate,
		Completed: false,
	}

	*todoList = append(*todoList, task)
}

func (todoList *TodoList) deleteTask(number int) error {
	index := number - 1

	if err := todoList.checkBounds(index); err != nil {
		return err
	}

	*todoList = append((*todoList)[:index], (*todoList)[index+1:]...)

	return nil
}

func (todoList *TodoList) markCompleted(number int) error {
	index := number - 1

	if err := todoList.checkBounds(index); err != nil {
		return err
	}

	isCompleted := (*todoList)[index].Completed

	if !isCompleted {
		(*todoList)[index].Completed = true
	}

	return nil
}

func (todoList *TodoList) checkBounds(index int) error {
	if index < 0 || index >= len(*todoList) {
		err := errors.New("index out of bounds")
		return err
	}
	return nil
}

func (todoList *TodoList) printTodoList() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintf(w, "Title\tPriority\tDue Date\tCompleted\n")
	for _, task := range *todoList {
		completed := "✅"
		if !task.Completed {
			completed = "❌"
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n",
			task.Title,
			task.Priority,
			task.DueDate.Format("2006-01-02"),
			completed,
		)
	}
	fmt.Println()
}

func (todoList *TodoList) saveData(filename string) error {
	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadTodoList(filename string) (*TodoList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &TodoList{}, nil
		}
		return nil, err
	}

	var todoList TodoList
	err = json.Unmarshal(data, &todoList)
	return &todoList, err
}

func (p Priority) String() string {
	return [...]string{"Low", "Medium", "High"}[p]
}

func ParsePriority(s string) Priority {
	switch strings.ToLower(s) {
	case "low":
		return Low
	case "medium":
		return Medium
	case "high":
		return High
	default:
		return Low
	}
}
