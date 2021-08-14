package main

import (
	"fmt"
)

func main() {

	x := 10
	y := 5

	// Arithmetic Operator

	zAdd := x + y
	zSubs := x - y
	zMult := x * y
	zDiv := x / y
	zMod := x % y

	fmt.Println(zAdd)
	fmt.Println(zSubs)
	fmt.Println(zMult)
	fmt.Println(zDiv)
	fmt.Println(zMod)

	// Relational Operator => returnnya Boolean

	i := 10
	j := 5

	fmt.Println(i == j)
	fmt.Println(i != j)
	fmt.Println(i >= j)
	fmt.Println(i <= j)

	// Logical Operator

	a := true
	b := false

	// AND &&
	fmt.Println(a && b)

	// OR ||
	fmt.Println(a || b)

	// NOT !
	fmt.Println(!b)
}
