package main

import (
	"fmt"
)

type Int int

type Person struct {
	Name    string
	Surname string
	Age     string
	City    string
}

func personNameAndSurname(p *Person) (string, string, error) {
	return p.Name, p.Surname, fmt.Errorf("error: %s", "Name and Surname not found")
}

func (p *Person) printHello() {
	fmt.Println(p)
}

func tutorial(ok bool) {
	p := Person{
		Name: func() string {
			if ok {
				return "John"
			}
			return "Jane"
		}(),
		Surname: func() string {
			if ok {
				fmt.Println("John")
				return "Wick"
			}
			fmt.Println("Jane")
			return "Doe"
		}(),
	}
	setName := func() {
		p.Name = "John"
	}

	setName()
	fmt.Printf("His name: %s\n", p.Name)

	//johnWick := Person{
	//	Name:    "John",
	//	Surname: "Wick",
	//}
	//fmt.Println(johnWick)
	//add(&johnWick,
	//	"John",
	//	"Jane",
	//	"Jack",
	//	"Jill",
	//	"James",
	//	"Jenny",
	//)
	//fmt.Println(johnWick)

}

// main - Entry point of the program
func main() {
	tutorial(true)

	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(slice)
	//
	//slice2 := []int{11, 12, 13, 14, 15}
	//
	//slice3 := append(slice, slice2...)
	//fmt.Println(slice3)
	//
	//slice5 := make([]int, 5, 10)
	//fmt.Println(slice5)
	//fmt.Println(len(slice5))
	//fmt.Println(cap(slice5))

	// map
	// Not comparable -  Slices, maps, channels, func
	// m := map[int][]int{}
	// m[1] = append(m[1], []int{1, 2, 3}...)
	// value, ok := m[1]
	// if !ok {
	// 	fmt.Println("key not found")
	// 	return
	// }
	// fmt.Println(value)

	//slice := []int{1, 2, 3, 4, 5, 65, 6, 7, 8, 9, 10}
	// for i := 0; i < len(slice); i++ {
	// 	if i == 9 {
	// 		break
	// 	}
	// 	fmt.Printf("%d ", slice[i])
	// }

	//johnWick := Person{
	//	Name:    "John",
	//	Surname: "Wick",
	//	Age:     "40",
	//	City:    "New York",
	//}
	//
	//fmt.Println(johnWick.GetFullName())

	//for k, v := range johnWick {
	//	fmt.Printf("Key: %s -----> Value: %s\n", k, v)
	//}
}
