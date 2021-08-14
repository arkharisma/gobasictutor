package main

import (
	"fmt"
)

// Slice is a dynamic Array, it can increase its size not like Array which has static size

func main() {

	mySlice := []int{10, 20, 30, 40, 50, 60}

	for i, v := range mySlice {
		fmt.Printf("Index: %d, Value: %d.\n", i, v)
	}

	mySliceStr := []string{"John", "Bobby", "David", "Samir"}

	mySliceStr = append(mySliceStr, "Arka")

	for _, v := range mySliceStr {
		fmt.Printf("Value: %s \n", v)
	}
}
