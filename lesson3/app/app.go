package app

import (
	"encoding/json"
	"fmt"
)

type App struct {
	SomeField string `json:"someField"`
	someField string `json:"some_filed"`
}

type myInt int

func New() *App {
	a := &App{
		SomeField: "a",
		someField: "b",
	}
	jsonApp, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(jsonApp))
	return a
}
