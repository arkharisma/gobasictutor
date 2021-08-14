package main

import (
	"fmt"
)

const A string = "Ini adalah Constant"

func main() {
	fmt.Println(A)

	const X int = 10

	fmt.Println(X)

	z := 9 / X

	fmt.Println(z)
}
