package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var myString = "Hello Arka"

	fmt.Println(myString)

	fmt.Println(len(myString))

	// for i := 0; i < len(myString); i++ {
	// 	fmt.Println(myString[i])
	// }

	myUpperString := strings.ToUpper(myString)
	fmt.Println(myUpperString)

	myLowerString := strings.ToLower(myString)
	fmt.Println(myLowerString)

	fmt.Println(strings.Contains(myString, "Ark")) // Ark != ark

	resultSplit := strings.Split(myString, "")

	// for i := 0; i < len(resultSplit); i++ {
	// 	fmt.Println(resultSplit[i])
	// }

	for _, v := range resultSplit {
		fmt.Println(v)
	}

	// Konversi String to Integer
	var myString2 = "10"

	resultConvInt, err := strconv.Atoi(myString2)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resultConvInt)
	}

	// Konversi Integer to String
	resultConvStr := strconv.Itoa(100)
	fmt.Println(resultConvStr)
}
