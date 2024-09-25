package main

import (
	"fmt"
	"strings"
)

const (
	availableTasks = 100
	low            = iota - 1
	medium
	high
)

func main() {

	var (
		projectStatus = "in progress"
		createdTasks  = 90
		isCompleted   = false
	)

	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is: 2024-09-18 00:00:00")
	fmt.Println("Project: Task Management System")
	fmt.Println()

	fmt.Printf("Current project status: %s\n", strings.ToUpper(projectStatus))

	fmt.Printf("Tasks completed: %d out of %d\n", availableTasks-createdTasks, availableTasks)
	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", low, medium, high)
	fmt.Printf("Is the project completed? %v\n", isCompleted)

}
