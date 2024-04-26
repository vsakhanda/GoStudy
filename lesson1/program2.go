package main

import (
	"fmt"
)

type Task struct {
	id     int
	title  string
	isDone bool
}

var tasks = make(map[int]Task)
var nextID = 1

func main() {

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
			addNewTask()
		case 2:
			removeTask()
		case 3:
			showTasks()
		case 4:
			completeTask()
		case 5:
			fmt.Println("Your choice is exit program.")
			fmt.Println("See you next time. Thanks for using program")
			return
			//default:
			//	fmt.Println("Please select Action from list")

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

func addNewTask() {
	fmt.Println("1. Add new task")
	var title string
	fmt.Print("Add title: ")
	_, _ = fmt.Scanln(&title)
	tasks[nextID] = Task{id: nextID, title: title, isDone: false}
	fmt.Printf("You added task # %v %v \n", nextID, title)
	nextID++
}

func removeTask() {
	fmt.Printf("Remove task action\n")

	if isEmpty(tasks) {
		fmt.Println("Empty list add new task!")
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
}

func showTasks() {
	if isEmpty(tasks) {
		fmt.Println("Empty list add new task!")
		return
	}
	displayTasks(tasks)
}

func completeTask() {
	fmt.Printf("Complete task\n")
	if isEmpty(tasks) {
		fmt.Println("Empty list add new task!")
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
}

//```go
//func main() {
//    reader := bufio.NewReader(os.Stdin)
//    fmt.Println("Введіть назву завдання:")
//    title, _ := reader.ReadString('\n')
//    fmt.Println("Ви ввели:", title)
//}
