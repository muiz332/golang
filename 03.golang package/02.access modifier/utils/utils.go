package utils

import "fmt"

var Nama string = "muiz" // bisa diakses dari luar
var umur uint8 = 10      // tidak bisa diakses dari luar

func SayHello() { // bisa diaksess
	fmt.Println("hello")
}

func sayHi() { // tidak bisa diakses
	fmt.Println("hi")
}
