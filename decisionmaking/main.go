package main

import (
	"fmt"
)

func main() {

	// If, Else If, Else
	if true {
		fmt.Println("Benar")
	} else if true && true {
		fmt.Println("Else If")
	} else {
		fmt.Println("Salah")
	}

	// Switch Case

	// Expression Switch
	x := 100

	switch x {
	case 60:
		fmt.Println(60)
	case 100:
		fmt.Println(100)
	default:
		fmt.Println("None of it")
	}

	switch {
	case x == 60:
		fmt.Println(60)
	case x == 100:
		fmt.Println(100)
	default:
		fmt.Println("None of it")
	}

	switch x := 100; {
	case x < 100:
		fmt.Println(60)
	case x > 100:
		fmt.Println(100)
	default:
		fmt.Println("None of it")
	}

	// Type Switch

	var y interface{}
	y = 5.9

	switch y.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float64")
	case float32:
		fmt.Println("Float32")
	default:
		fmt.Println("None of it")
	}
}
