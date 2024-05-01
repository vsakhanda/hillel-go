package main

import (
	"fmt"
	"time"
)

type Employee struct {
	id      int
	name    string
	surname string
}

var employees = make(map[int]Employee)
var id = 1

type Schedule struct {
	nextID           int
	employeeId       int
	date             time.Weekday
	startTime        time.Time
	endTime          time.Time
	totalWorkedHours time.Duration
}

var works = make(map[int]Schedule)
var nextID = 1

func main() {

	fmt.Println("Start application")
	fmt.Println("##################")
	fmt.Println("List of actions:")
	fmt.Println("1. Add new Employee")
	fmt.Println("2. Change employee data")
	fmt.Println("3. Show employee list")
	fmt.Println("4. Add new work information")
	fmt.Println("5. Show work information for employee")
	fmt.Println("6. Show work information for date")
	fmt.Println("7. Show total working hours for all employees")
	fmt.Println("8. Exit")
	fmt.Println("##################")

	for {
		var selectAction int
		fmt.Print("Select action: ")
		fmt.Scanln(&selectAction)

		switch selectAction {
		case 1:
			fmt.Println("1. Add new Employee")
			addEmployee()
		case 2:
			fmt.Println("2. Show employee list")
			showEmployeesList()
		case 3:
			fmt.Println("3. Change employee data")
			changeEmployeeData()
		case 4:
			fmt.Println("4. Add new work information")
		case 5:
			fmt.Println("5. Show work information for employee")
		case 6:
			fmt.Println("6. Show work information for date")
		case 7:
			fmt.Println("7. Show total working hours for all employees")
		case 8:
			fmt.Println("Your choice is to exit program.")
			fmt.Println("See you next time. Thanks!")
			return
			//default:
			//	fmt.Println("Please select Action from list")

		}
	}

}

// Functions:

// Введення ім'я співробітника.
func addEmployee() {
	var name string
	var surname string

	fmt.Print("Add your name: ")
	_, _ = fmt.Scanln(&name)
	fmt.Print("Add your surname: ")
	_, _ = fmt.Scanln(&surname)
	employees[id] = Employee{id: id, name: name, surname: surname}
	fmt.Printf("Employee added # %v %v %v \n", id, name, surname)
	id++
}

func showEmployeesList() {
	fmt.Println("Employees: #, Name, Surname")
	for id, employee := range employees {
		fmt.Printf("# %v %v  %v \n", id, employee.name, employee.surname)
	}
}

func isNoEmployee(employees map[int]Employee) bool {
	return len(employees) == 0
}

func changeEmployeeData() {
	//fmt.Println("3. Change employee data")
	isNoEmployee(employees)
	if isNoEmployee(employees) {
		fmt.Println("Empty list add new task!")
	}
	showEmployeesList()
	var id int
	fmt.Print("Enter task ID for complete task: ")
	fmt.Scanln(&id)
	if employee, exists := employees[id]; exists {
		var newName string
		var newSurmane string
		fmt.Print("Enter new name: ")
		fmt.Scanln(&newName)
		fmt.Print("Enter new surname: ")
		fmt.Scanln(&newSurmane)
		employee.name = newName
		employee.surname = newSurmane
		employees[id] = employee
		fmt.Printf("# %v %v  %v \n", id, employee.name, employee.surname)
		fmt.Println("Employee data up to date.")
	} else {
		fmt.Printf("Incorrect task # %v id.", id)
	}

}

//Введення часу приходу та відходу в форматі годин та хвилин. - ідентифікація виконується через id користувача та вивделення даних для підтвердження

func inputWorktime() {
	fmt.Println("Entre your id")

}

//Розрахунок відпрацьованих годин за день
//workDurationCalculation
//Збереження та вивід загальної кількості відпрацьованих годин за тиждень.
//collectTotalWorkingHours
// вивід поточної інформації
//showCurrentWorkingHours
// збереження інформації про підпрацьовані години для співробітника - може бути реалізована як части функції введення даних

// перевірка на наявність співробітнка

//

//
//Створіть програму, яка дозволяє вводити час приходу та відходу співробітників, а потім розраховує відпрацьовані години за день. Програма повинна також зберігати загальну кількість відпрацьованих годин за тиждень.
//Функціональні вимоги:
//Введення ім'я співробітника.
//Введення часу приходу та відходу в форматі годин та хвилин.
//Розрахунок відпрацьованих годин за день.
//Збереження та вивід загальної кількості відпрацьованих годин за тиждень.
//Обмеження:
//Дозволено використання наступних пакетів
//fmt
//strings
//strconv
//time
//bufio/os - тільки для вводу тексту, опціонально
//Вивантаження дз:
//Склонувати репозиторій hillel-go, який закріпленний до уроку, у свій аккаунт, запушити вашу роботу в ваш клонований репозиторій та створити pull request у моему репозиторії. Надати лінк на pull request.
