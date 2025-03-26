package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string // Slice to store tasks

func addTask() {
	fmt.Print("Enter task: ")
	reader := bufio.NewReader(os.Stdin)
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task) // Remove newline character
	tasks = append(tasks, task)
	fmt.Println("Task added!")
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
		return
	}
	fmt.Println("\nYour To-Do List:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func main() {
	for {
		fmt.Println("\n1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			viewTasks()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option! Try again.")
		}
	}
}
