package main

import (
	"fmt"
	"time"
)

type TimeSheet struct {
	tsId      int
	EmpID     int
	name      string
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
var IdEmp = 1

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
			showAll()
		case 2:
			fmt.Println("2. Add information")
			addWorkData()
		case 3:
			fmt.Println("3. Show total working hours for all employees")
			showTime()
		case 4:
			showEmployeeByName()
		case 5:
			fmt.Println("Your choice is to exit program.")
			fmt.Println("See you next time. Thanks!")
			return
		}
	}
}

func showAll() {
	fmt.Println("List of actions:")
	fmt.Println("1. Show all actions")
	fmt.Println("2. Add information")
	fmt.Println("3. Show total working hours for all employees")
	fmt.Println("4. Show information for employee")
	fmt.Println("5. Exit")
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
	var employeeID int
	var name string
	var dayOfWeek int
	var startTimeStr string
	var endTimeStr string
	var workedHours time.Duration

	fmt.Print("Enter your id: ")
	_, _ = fmt.Scanln(&employeeID)

	fmt.Print("Add your name: ")
	_, _ = fmt.Scanln(&name)

	fmt.Print("Enter start of work time: (HH:MM format)")
	fmt.Scan(&startTimeStr)
	startTime, err := time.Parse("15:04", startTimeStr)
	if err != nil {
		fmt.Println("Invalid time format. Please use HH:MM format.")
		return
	}

	fmt.Print("Enter end of work (HH:MM format): ")
	fmt.Scan(&endTimeStr)
	endTime, err := time.Parse("15:04", endTimeStr)
	if err != nil {
		fmt.Println("Invalid time format. Please use HH:MM format.")
		return
	}

	workedHours = endTime.Sub(startTime)

	fmt.Print("Enter day number: ")
	fmt.Scan(&dayOfWeek)
	if dayOfWeek < 1 || dayOfWeek > 7 {
		fmt.Println("Некоректний ввід.")
		return
	}
	switch dayOfWeek {
	case 1:
		timeSheets[tsId] = TimeSheet{tsId: tsId, EmpID: employeeID, name: name, monday: workedHours}
	case 2:
		timeSheets[tsId] = TimeSheet{tsId: tsId, EmpID: employeeID, name: name, tuesday: workedHours}
	case 3:
		timeSheets[tsId] = TimeSheet{tsId: tsId, EmpID: employeeID, name: name, wednesday: workedHours}
	case 4:
		timeSheets[tsId] = TimeSheet{tsId: tsId, EmpID: employeeID, name: name, thursday: workedHours}
	case 5:
		timeSheets[tsId] = TimeSheet{tsId: tsId, EmpID: employeeID, name: name, friday: workedHours}
	case 6:
		timeSheets[tsId] = TimeSheet{tsId: tsId, EmpID: employeeID, name: name, saturday: workedHours}
	case 7:
		timeSheets[tsId] = TimeSheet{tsId: tsId, EmpID: employeeID, name: name, sunday: workedHours}
	default:
		fmt.Println("Unknown day")
		return
	}

	fmt.Printf("Worked hours %v mon %v tue %v wed %v thu %v fri %v sat %v sun %v %v total \n",
		timeSheets[tsId].EmpID,
		timeSheets[tsId].name,
		timeSheets[tsId].monday,
		timeSheets[tsId].tuesday,
		timeSheets[tsId].wednesday,
		timeSheets[tsId].thursday,
		timeSheets[tsId].friday,
		timeSheets[tsId].saturday,
		timeSheets[tsId].sunday,
		timeSheets[tsId].totalTime,
	)
	tsId++
}

func showEmployeeByName() {
	fmt.Print("Enter user name for additional information: ")
	fmt.Scanln(&IdEmp)
	var timeSheets, exists = timeSheets[IdEmp]
	if exists {
		fmt.Printf("Worked by # %v employee %v mon %v tue %v wed %v thu %v fri %v sat %v sun %v %v total \n",
			timeSheets.EmpID,
			timeSheets.name,
			timeSheets.monday,
			timeSheets.tuesday,
			timeSheets.wednesday,
			timeSheets.thursday,
			timeSheets.friday,
			timeSheets.saturday,
			timeSheets.sunday,
			timeSheets.totalTime,
		)
	}
}
func showTime() {
	//if isEmptyList(timeSheets) {
	//	fmt.Println("Add first work in table")
	//}
	fmt.Println("Schedule Table: #, Employee Id, StartTime, EndTime, WorkingHours")
	for _, timeShow := range timeSheets {
		fmt.Printf("Worked hours # mon %v tue %v wed %v thu %v fri %v sat %v sun %v %v total \n",
			timeShow.monday,
			timeShow.tuesday,
			timeShow.wednesday,
			timeShow.thursday,
			timeShow.friday,
			timeShow.saturday,
			timeShow.sunday,
			timeShow.totalTime,
		)
	}
}

//func isEmptyList(works map[int]Schedule) bool {
//	return len(works) == 0
//}
