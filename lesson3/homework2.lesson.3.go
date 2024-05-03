package main

import (
	"fmt"
	"time"
)

type TimeSheet struct {
	tsId      int
	Empolyee  string
	monday    time.Duration
	tuesday   time.Duration
	wednesday time.Duration
	thursday  time.Duration
	friday    time.Duration
	saturday  time.Duration
	sunday    time.Duration
	totalTime time.Duration
}

var timeSheets = make(map[int]TimeSheet)
var tsId = 1

func main() {

	fmt.Println("Start application")
	fmt.Println("##################")
	showAll()
	fmt.Println("##################")

	for {
		var selectAction int
		fmt.Print("#### To show list of actions input 1:\n")
		fmt.Print("Select action: ")
		fmt.Scanln(&selectAction)
		switch selectAction {
		case 1:
			showAllActions()
		case 2:
			fmt.Println("2. Add information")
			addWorkData()
		case 3:
			fmt.Println("3. Show total working hours for all employees")
			showTimeSheet()
		case 4:
			fmt.Println("Your choice is to exit program.")
			fmt.Println("See you next time. Thanks!")
			return
		}
	}
}

// Functions:
func showAll() {
	fmt.Println("List of actions:")
	fmt.Println("1. Show all actions")
	fmt.Println("2. Add information")
	fmt.Println("3. Show total working hours for all employees")
	fmt.Println("4. Exit")

}

func daysOfWeek() {
	fmt.Println("Select day of work")
	fmt.Println("1. Monday")
	fmt.Println("2. Tuesday")
	fmt.Println("3. Wednesday")
	fmt.Println("4. Thursday")
	fmt.Println("5. Friday")
	fmt.Println("6. Saturday")
	fmt.Println("7. Sunday")
}

func addWorkData() {
	var name string
	var dayOfWeek int
	var startTimeStr string
	var endTimeStr string
	var workedHours time.Duration

	fmt.Print("Add your name: ")
	_, _ = fmt.Scanln(&name)

	fmt.Print("Start of work: ")
	fmt.Scan(&startTimeStr)
	startTime, err := time.Parse("15:04", startTimeStr)
	if err != nil {
		fmt.Println("Invalid time format. Please use HH:MM format.")
		return
	}
	fmt.Print("Enter end of work (YYYY:MM:DD:HH:MM format): ")
	fmt.Scan(&endTimeStr)
	endTime, err := time.Parse("15:04", endTimeStr)
	if err != nil {
		fmt.Println("Invalid time format. Please use HH:MM format.")
		return
	}

	workedHours = endTime.Sub(startTime)

	// Зчитування дня тижня
	fmt.Print("Enter day number: ")
	fmt.Scan(&dayOfWeek)
	if dayOfWeek < 1 || dayOfWeek > 7 {
		fmt.Println("Некоректний ввід.")
		return
	}

	// Вибір поля, в яке слід записати тривалість роботи
	switch dayOfWeek {
	case 1:
		dayOfWeek = timeSheets[dayOfWeek].monday
	case 2:
		dayOfWeek = timeSheets[id].tuesday
	case 3:
		dayOfWeek = timeSheets[id].wednesday
	case 4:
		dayOfWeek = timeSheets[id].thursday
	case 5:
		dayOfWeek = timeSheets[id].friday
	case 6:
		dayField = timeSheets[id].saturday
	case 7:
		dayField = timeSheets[id].sunday
	default:
		fmt.Println("Unknown day")
		return
	}

	// Запис тривалості роботи в поле та оновлення totalTime
	dayField = workedHours
	timeSheets[id].totalTime += workedHours

	// Виведення результату
	fmt.Printf("Worked hours # mon %v tue %v wed %v thu %v fri %v sat %v sun %v %v total \n",
		timeSheets[id].monday,
		timeSheets[id].tuesday,
		timeSheets[id].wednesday,
		timeSheets[id].thursday,
		timeSheets[id].friday,
		timeSheets[id].saturday,
		timeSheets[id].sunday,
		timeSheets[id].totalTime,
	)
	tsId++
}

func showEmployeeByID1() {
	fmt.Print("Enter user id for additional information: ")
	fmt.Scanln(&id)
	employee, exists := employees[id]
	if exists {
		fmt.Printf("ID: %d, Name: %s, Surname: %s \n", employee.id, employee.name, employee.surname)
	} else {
		fmt.Printf("Incorrect employee # %v id.", id)
	}
}
func showTimeSheet() {
	if isEmptyList(Works) {
		fmt.Println("Add first work in table")
	}
	fmt.Println("Schedule Table: #, Employee Id, StartTime, EndTime, WorkingHours")
	for _, schedule := range Works {
		fmt.Printf("# %v user %v  startTime %v end time %v work hours %v \n", schedule.nextID, schedule.employeeId, schedule.startTime.Format("15:04"), schedule.endTime.Format("15:04"), schedule.totalWorkedHours)
	}
}
func isEmptyList(works map[int]Schedule) bool {
	return len(works) == 0
}
