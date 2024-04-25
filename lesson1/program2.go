package main

import (
	"fmt"
)

type Task struct {
	id     int
	title  string
	isDone bool
}

func main() {

	tasks := make(map[int]Task)
	nextID := 1

	fmt.Println("##################")
	fmt.Println("List of actions:")
	fmt.Println("1. Add new task")
	fmt.Println("2. Remove task")
	fmt.Println("3. Show task list")
	fmt.Println("4. Complete task")
	fmt.Println("5. Exit")
	fmt.Println("##################")

	for {

		var selectAction (int)
		fmt.Print("Select Action: ")
		fmt.Scanln(&selectAction)

		switch selectAction {
		case 1:
			fmt.Println("1. Add new task")
			var title string
			fmt.Print("Add title: ")
			fmt.Scanln(&title)
			tasks[nextID] = Task{id: nextID, title: title, isDone: false}
			nextID++
		case 2:
			fmt.Printf("Remove task action\n")

			if isEmpty(tasks) {
				fmt.Println("Empty list add new task!")
				continue
			}
			displayTasks(tasks)
			var taskID int
			fmt.Print("Enter select ID for delete task: ")
			fmt.Scanln(&taskID)

			if task, exists := tasks[taskID]; exists {
				tasks[taskID] = task
				fmt.Printf("Task %v deleted successfully!\n", taskID)
				delete(tasks, taskID)
			} else {
				fmt.Println("Incorrect task ID!")
			}
		case 3:
			if isEmpty(tasks) {
				fmt.Println("Empty list add new task!")
				continue
			}
			displayTasks(tasks)
		case 4:
			fmt.Printf("Complete task\n")
			if isEmpty(tasks) {
				fmt.Println("Empty list add new task!")
				continue
			}
			displayTasks(tasks)
			var taskID int
			fmt.Print("Enter task ID for complete task: ")
			fmt.Scanln(&taskID)
			if task, exists := tasks[taskID]; exists {
				task.isDone = !task.isDone
				tasks[taskID] = task
				fmt.Println("Task completed successfully!")
			} else {
				fmt.Println("Incorrect task ID!")
			}
		case 5:
			fmt.Println("Your choice is exit program.")
			fmt.Println("See you next time. Thanks for using program")
			return
		default:
			fmt.Println("Please select Action from list")

		}
	}
}

// перевірка на пустий список
func isEmpty(tasks map[int]Task) bool {
	return len(tasks) == 0
}

// відображення списку задач
func displayTasks(tasks map[int]Task) {
	fmt.Println("Task list:")
	for id, task := range tasks {
		fmt.Printf("%d. %s (Completed: %t)\n", id, task.title, task.isDone)
	}
}
