package main

import (
	"fmt"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks = make(map[int]Task)
var idCounter int

// Нове завдання (функція)
func addTask(title string) {
	tasks[idCounter] = Task{ID: idCounter, Title: title, Completed: false}
	idCounter++
}

// Мінус завдання за ID (функція)
func deleteTask(id int) {
	_, exists := tasks[id]
	if exists {
		delete(tasks, id)
		return
	}
	fmt.Println("Завдання з таким ID не існує!")
}

// Перегляд усіх завдань (метод)
func (t Task) viewTasks() {
	fmt.Println("Усі завдання:")
	for _, task := range tasks {
		status := "не завершене"
		if task.Completed {
			status = "завершене"
		}
		fmt.Printf("ID: %d, Завдання: %s, Статус: %s\n", task.ID, task.Title, status)
	}
}

// Завершененя завдання (метод)
func (t *Task) markAsCompleted() {
	t.Completed = true
}

func main() {
	fmt.Println("Введіть що робити:")
	fmt.Println("1 - додати")
	fmt.Println("2 - видалити")
	////

	for {
		var action int
		fmt.Scanln(&action)
		switch action {
		case 1:
			fmt.Println("Введіть назву завдання:")
			var title string
			fmt.Scanln(&title)
			addTask(title)
		case 2:
			fmt.Println("Введіть ID завдання:")
			var id int
			fmt.Scanln(&id)
			deleteTask(id)
		}
	}
}
