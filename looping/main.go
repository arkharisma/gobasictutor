package main

import (
	"fmt"
)

func main() {

	// For Looping
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	x := 10
	var y int

	for y < x {
		fmt.Println(y)
		y++
	}

	// Go tidak ada while, untuk penggantinya bisa dibuat begini
	z := 0
	for {
		fmt.Println(z)
		z++
		if z == 10 {
			break
		}
	}

}
