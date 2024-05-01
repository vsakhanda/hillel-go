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
	startTime        time.Time
	endTime          time.Time
	totalWorkedHours time.Duration
}

var Works = make(map[int]Schedule)
var nextID = 1

func main() {

	fmt.Println("Start application")
	fmt.Println("##################")
	showAllActions()
	fmt.Println("##################")

	for {
		var selectAction int

		fmt.Print("To show list of actions input 1:\n")
		fmt.Print("Select action: ")
		fmt.Scanln(&selectAction)

		switch selectAction {
		case 1:
			showAllActions()
		case 2:
			fmt.Println("2. Add new Employee")
			addEmployee()
		case 3:
			fmt.Println("3. Show employee list")
			showEmployeesList()
		case 4:
			fmt.Println("4. Change employee data")
			changeEmployeeData()
		case 5:
			fmt.Println("5. Show work information for employee")
			showEmployeeByID()
		case 6:
			fmt.Println("6. Add new work information")
			inputWorkTime()
		case 7:
			fmt.Println("7. Show work information for week")
			workDurationCalculation()
		case 8:
			showScheduleTable()
		case 9:
			fmt.Println("Your choice is to exit program.")
			fmt.Println("See you next time. Thanks!")
			return
			//default:
			//	fmt.Println("Please select Action from list")

		}
	}

}

// Functions:
func showAllActions() {
	fmt.Println("List of actions:")
	fmt.Println("1. Show all actions")
	fmt.Println("2. Add new Employee")
	fmt.Println("3. Change employee data")
	fmt.Println("4. Show employee list")
	fmt.Println("5. Show work information for employee")
	fmt.Println("6. Add new work information")
	fmt.Println("7. Show work information for date")
	fmt.Println("8. Show total working hours for all employees")
	fmt.Println("9. Exit")

}

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
	if isNoEmployee(employees) {
		fmt.Println("Empty list add first employee! Action #2")
	}
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
		fmt.Println("Empty list of users!")
	}
	//showEmployeesList()
	var id int
	fmt.Print("Enter your id: ")
	fmt.Scanln(&id)
	if employee, exists := employees[id]; exists {
		var newName string
		var newSurname string
		fmt.Print("Enter new Name: ")
		fmt.Scanln(&newName)
		fmt.Print("Enter new Surname: ")
		fmt.Scanln(&newSurname)
		employee.name = newName
		employee.surname = newSurname
		employees[id] = employee
		fmt.Printf("# %v %v  %v \n", id, employee.name, employee.surname)
		fmt.Println("Employee data up to date.")
	} else {
		fmt.Printf("Incorrect employee's # %v id. \n", id)
	}

}

func showEmployeeByID() {
	fmt.Print("Enter user id for additional information: ")
	fmt.Scanln(&id)
	employee, exists := employees[id]
	if exists {
		fmt.Printf("ID: %d, Name: %s, Surname: %s \n", employee.id, employee.name, employee.surname)
	} else {
		fmt.Printf("Incorrect employee # %v id.", id)
	}
}

func inputWorkTime() {
	var employeeId int
	var startTimeStr string
	var endTimeStr string
	var WorkedHours time.Duration

	fmt.Print("Enter your id: ")
	fmt.Scanln(&employeeId)

	//fmt.Print("Enter start of work: ")
	//_, _ = fmt.Scanln(&startTime)
	fmt.Print("Enter start of work (YYYY:MM:DD:HH:MM format): ")
	fmt.Scan(&startTimeStr)
	startTime, err := time.Parse("2006:01:02:15:04", startTimeStr)
	if err != nil {
		fmt.Println("Invalid time format. Please use HH:MM format.")
		return
	}
	fmt.Print("Enter end of work (YYYY:MM:DD:HH:MM format): ")
	fmt.Scan(&endTimeStr)
	endTime, _ := time.Parse("2006:01:02:15:04", endTimeStr)
	if err != nil {
		fmt.Println("Invalid time format. Please use HH:MM format.")
		return
	}
	WorkedHours = endTime.Sub(startTime)
	Works[nextID] = Schedule{nextID: nextID, employeeId: employeeId, startTime: startTime, endTime: endTime, totalWorkedHours: WorkedHours}
	fmt.Printf(" Employee %v worked for %v \n ", employeeId, WorkedHours)
	fmt.Printf("Enter line: #%v Employee id %v worked from %v to %v for duration of %v Hours and minutes \n", nextID, employeeId, startTime.Format("15:04"), endTime.Format("15:04"), WorkedHours)
	nextID++
}

func workDurationCalculation() {

	totalHours := make(map[int]time.Duration)

	for _, work := range Works {
		totalHours[work.employeeId] += work.totalWorkedHours
	}

	fmt.Println("Total hours for user:")
	for employeeID, hours := range totalHours {
		fmt.Printf("Employee #%d: %v\n", employeeID, hours)
	}
}

func showScheduleTable() {
	if isEmpty(Works) {
		fmt.Println("Add first work in table")
	}
	fmt.Println("Schedule Table: #, Employee Id, StartTime, EndTime, WorkingHours")
	for _, schedule := range Works {
		fmt.Printf("# %v user %v  startTime %v end time %v work hours %v \n", schedule.nextID, schedule.employeeId, schedule.startTime.Format("15:04"), schedule.endTime.Format("15:04"), schedule.totalWorkedHours)
	}
}
func isEmpty(works map[int]Schedule) bool {
	return len(works) == 0
}

/*	fmt.Println("############### - Full list ")
	fmt.Println("Employees: #, Name, Surname")
	for id, employee := range employees {
		if id == id {
			// Трошки говнокодіку - але перевірив як цікаво працює =)
			fmt.Printf("# %v %v %v\n", id, employee.name, employee.surname)
			//
			break
		}
	}*/
