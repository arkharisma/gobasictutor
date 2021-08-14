package main

import (
	"fmt"
)

func main() {

	p := &Person{ID: 1, Name: "Arka", Age: 23}

	fmt.Printf("ID: %d, Name: %s, Age: %d.\n", p.GetID(), p.GetName(), p.GetAge())
	p.ChangeName("Arka Kharisma")
	fmt.Printf("ID: %d, Name: %s, Age: %d.\n", p.GetID(), p.GetName(), p.GetAge())
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func (p *Person) GetID() int {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}
