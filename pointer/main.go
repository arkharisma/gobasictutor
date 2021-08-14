package main

import (
	"fmt"
)

func main() {

	var hello string = "hello"
	var strPtr *string

	// Deklarasi nilai strPtr yaitu nilai alamat dari variable hello
	strPtr = &hello

	// Alamat dari variabel hello
	fmt.Println(strPtr)
	// Nilai dari alamat variable hello
	fmt.Println(*strPtr)

	change(hello)
	fmt.Println(hello)

	changeByPtr(&hello)
	fmt.Println(hello)
}

// Fungsi yang tidak menggunakan call by reference, sehingga perubahan yang ada hanya local di function
func change(v string) {
	v = v + " Golang"
}

// Fungsi menggunakan call by reference, karena nilai yang diubah itu langsung menunjuk ke alamat si variable
func changeByPtr(v *string) {
	*v = *v + " Golang"
}
