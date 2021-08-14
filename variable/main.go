package main

import (
	"fmt"
)

func main() {

	// static type declaration, harus dideklarasikan tipe datanya
	var x int
	x = 10

	var y float64
	y = 5.5

	fmt.Printf("nilai x adalah %d \n", x)
	fmt.Printf("nilai y adalah %.1f \n", y)

	fmt.Printf("x memiliki tipe data: %T \n", x)
	fmt.Printf("y memiliki tipe data: %T \n", y)

	// dynamic type declaration, tidak perlu dideklarasikan tipe data, dan Go sudah otomatis tahu
	z := "arka"
	i := 50
	fmt.Println(z)
	fmt.Println(i)

	a := 5
	b := 10

	c := a + b
	fmt.Println(c)
}
