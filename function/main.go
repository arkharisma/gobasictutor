package main

import (
	"fmt"
)

func main() {

	x := 5
	y := 10
	z := add(x, y)
	fmt.Println(z)

	// Concantenate String Return Function
	fmt.Println("Concatenate String Return Function")
	name := "Arka"
	result := hello(name)

	fmt.Println(result)

	// Multiple Return Function
	fmt.Println("Multiple Return Function")
	resultXSwap, resultYSwap := swap("Arka", "Kharisma")

	fmt.Println(resultXSwap, resultYSwap)

	// Function as Value
	substract := func(x, y int) int {
		return x - y
	}

	fmt.Println(substract(5, 3))

	// Closure (Fungsi dengan return function)
	nextValue := genNumber()

	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())

	arkaLove := love("Arka")

	fmt.Println(arkaLove("Alma"))
	fmt.Println(arkaLove("Cat"))

	// Function as Parameter
	f := func(v string) bool {
		return v == "Go"
	}

	resultF := match("Go", f)

	fmt.Println(resultF)

}

// Function as Parameter
func match(v string, f func(string) bool) bool {
	if f(v) {
		return true
	}
	return false
}

// Closure
func genNumber() func() int {
	// value x di deklarasikan pada waktu pemanggilan genNumber(), dan jika nextValue di invoke lagi, maka yang berjalan adalah function returnnya
	x := 0
	return func() int {
		x += 1
		return x
	}
}

// Closure
func love(name string) func(string) string {
	return func(things string) string {
		return fmt.Sprintf("%s love %s", name, things)
	}
}

func add(x int, y int) int {
	return x + y
}

func hello(z string) string {
	return fmt.Sprintf("Hello %s", z)
}

func swap(x string, y string) (string, string) {
	return y, x
}
