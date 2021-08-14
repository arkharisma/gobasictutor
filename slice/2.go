package main

import (
	"fmt"
)

func main() {
	slicePersons := []Person{Person{ID: 1, Name: "Arka", Age: 23}, Person{ID: 2, Name: "John", Age: 35}}

	for _, v := range slicePersons {
		fmt.Printf("ID: %d, Name: %s, Age: %d.\n", v.ID, v.Name, v.Age)
	}
}

type Person struct {
	ID   int
	Name string
	Age  int
}
