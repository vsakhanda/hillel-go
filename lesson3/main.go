package main

import (
	"fmt"
	"strings"
	"time"
)

type SomeStruct struct {
	Field1 string
	Field2 string
	Field3 string
}

func someRuneFunc() func(r rune) bool {
	return func(r rune) bool {
		return false
	}
}

func main() {
	// time, strconv, strings

	// 05.12.2024 02:52:33
	// 02.52.22 - 12:05:2024
	// 02:54:22 05.12.2024
	//t, err := time.Parse("15:04:05 02.01.2006", "02:54:22 05.12.2024")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(t.Compare(time.Now()))

	//var number int64 = 517983457256213
	//numberString := "100"

	//parsedNumber, err := strconv.Atoi(numberString)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Sprintf("%d", number)
	//numberString = strconv.Itoa(number)

	//numberString = strconv.FormatInt(number, 16)
	//fmt.Println(numberString)
	//
	//var someFloat float32 = 125.35161681
	//
	//floatString := strconv.FormatFloat(float64(someFloat), 'f', 10, 32)
	//fmt.Printf("%.2f%%", someFloat)
	//
	//boolean := strconv.FormatBool(true)
	//_ = boolean

	someStringWithSpaces := "Johnk Wick"
	someStringWithSpaces = strings.TrimLeft(someStringWithSpaces, "@")
	_ = someStringWithSpaces

	i := strings.Compare(someStringWithSpaces, "John Wic")
	_ = i

	someStringWithSpaces = strings.ToUpper(someStringWithSpaces)
	_ = someStringWithSpaces

	someStringWithSpaces = strings.ToLower(someStringWithSpaces)
	_ = someStringWithSpaces

	strings.Contains(someStringWithSpaces, "wick")

	//fmt.Println(strings.ContainsFunc(someStringWithSpaces, someRuneFunc()))

	fmt.Println(strings.HasSuffix(someStringWithSpaces, "ck"))

	//someStringWithSpaces = strings.Replace(someStringWithSpaces, "k", "a", -1)
	//_ = someStringWithSpaces

	address := "м. Київ, вул. Хрещатик, буд. 55"
	someSlice := strings.Split(address, ", ")
	_ = someSlice

	joinedString := strings.Join(someSlice, "|")
	_ = joinedString

	someStruct := SomeStruct{
		Field1: "John",
		Field2: "",
		Field3: "Wick",
	}

	builder := strings.Builder{}
	if someStruct.Field1 != "" {
		builder.WriteString(someStruct.Field1 + ", ")
	}
	if someStruct.Field2 != "" {
		builder.WriteString(someStruct.Field2 + ", ")
	}
	if someStruct.Field3 != "" {
		builder.WriteString(someStruct.Field3 + ", ")
	}
	str := builder.String()
	str = strings.TrimRight(str, ", ")
	str += "!"
	_ = str

	time.Sleep(5 * time.Second)
	fmt.Println("End of program")
}
