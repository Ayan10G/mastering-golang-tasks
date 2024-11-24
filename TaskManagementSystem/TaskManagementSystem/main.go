package main

import (
	"fmt"
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
		createdTasks  = 100
		complatedtask = 95
		isCompleted   = false
	)

	statikMelumatlar()

	if complatedtask > createdTasks {
		fmt.Println("There was an error on inputs...reseting the completed task")
		complatedtask = createdTasks

	} else if complatedtask == 0 {
		panic("Project is not started yet")
	}

	fmt.Printf("Tasks remaining %d out of %d\n", createdTasks-complatedtask, createdTasks)

	if createdTasks == complatedtask {
		projectStatus = "Done"
		isCompleted = true
	}

	fmt.Println("Current project status: ", projectStatus)

	switch {
	case complatedtask < 30:
		fmt.Println("Project is in the starting phase.")
	case complatedtask >= 30 && complatedtask <= 60:
		fmt.Println("Project is in the Midway phase.")
	case complatedtask > 60:
		fmt.Println("Project is in the Final phase")
	}

	fmt.Println("Task list:")

	for ayan := 1; ayan < 6; ayan++ {
		fmt.Printf("- Task %d \n", ayan)
	}

	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", low, medium, high)
	fmt.Printf("Is the project completed? %v\n", isCompleted)

}

func statikMelumatlar() {
	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is: 2024-09-18 00:00:00")
	fmt.Println("Project: Task Management System")
	fmt.Println()
}
