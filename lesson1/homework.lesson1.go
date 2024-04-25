package main

import "fmt"

//Завдання: Створіть програму для керування персональним списком завдань. Користувач може додавати нові завдання, видаляти існуючі, а також переглядати весь список.
//Обмеження: Використовувати тільки пакет "fmt"
//Вимоги:
//Структура "Завдання": Визначте структуру Task, яка включатиме поля Title (назва завдання) та Completed (чи завершено завдання).
//Мапа для зберігання завдань: Використайте мапу, де ключем буде унікальний ID завдання (наприклад, ціле число), а значенням — об'єкт Task.
//Функції для керування завданнями:
//Додавання нового завдання.
//Видалення завдання за ID.
//Перегляд усіх завдань.
//Позначення завдання як завершене.

type Task1 struct {
	id     int
	title  string
	isDone bool
}

func main() {

	tasks := make(map[int]Task1)
	nextID := 1

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Додавання нового завдання")
		fmt.Println("2. Видалення завдання за ID")
		fmt.Println("3. Перегляд усіх завдань")
		fmt.Println("4. Позначення завдання як завершене")
		fmt.Println("5. Вийти з програми")

		var choice int
		fmt.Print("Виберіть опцію: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title string
			fmt.Print("Введіть назву завдання: ")
			fmt.Scanln(&title)
			tasks[nextID] = Task1{id: nextID, title: title, isDone: false}
			nextID++
			fmt.Println("Завдання успішно додано!")
		case 2:
			if len(tasks) == 0 {
				fmt.Println("Список завдань порожній!")
				continue
			}
			fmt.Println("Список завдань:")
			for id, task := range tasks {
				fmt.Printf("%d. %s (Виконано: %t)\n", id, task.title, task.isDone)
			}
			var taskID int
			fmt.Print("Виберіть ID завдання для видалення: ")
			fmt.Scanln(&taskID)
			if _, exists := tasks[taskID]; !exists {
				fmt.Println("Неправильний ID завдання!")
				continue
			}
			delete(tasks, taskID)
			fmt.Println("Завдання успішно видалено!")
		case 3:
			if len(tasks) == 0 {
				fmt.Println("Список завдань порожній!")
				continue
			}
			fmt.Println("Список завдань:")
			for id, task := range tasks {
				fmt.Printf("%d. %s (Виконано: %t)\n", id, task.title, task.isDone)
			}
		case 4:
			if len(tasks) == 0 {
				fmt.Println("Список завдань порожній!")
				continue
			}
			fmt.Println("Список завдань:")
			for id, task := range tasks {
				fmt.Printf("%d. %s (Виконано: %t)\n", id, task.title, task.isDone)
			}
			var taskID int
			fmt.Print("Виберіть ID завдання для зміни стану: ")
			fmt.Scanln(&taskID)
			if task, exists := tasks[taskID]; exists {
				task.isDone = !task.isDone
				tasks[taskID] = task
				fmt.Println("Стан завдання успішно змінено!")
			} else {
				fmt.Println("Неправильний ID завдання!")
			}
		case 5:
			fmt.Println("Дякую за використання програми!")
			return
		default:
			fmt.Println("Неправильний вибір опції. Спробуйте ще раз.")
		}
	}

}
