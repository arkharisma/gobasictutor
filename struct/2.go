package main

import (
	"fmt"
)

func main() {
	person1 := Person{ID: 1, Name: "Arka", Age: 23}
	printPerson(person1)
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func printPerson(p Person) {
	fmt.Println(p.ID)
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
