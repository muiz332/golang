/*

- check type data


ada kalanya kita ingin mengecheck apakah type data yang diterima itu
berbentuk string,integer, float dan lain lain


kita bisa gunakan beberapa cara berikut ini
ada 3 cara untuk mengetahui type data


- menggunakan %T printf
- menggunakan reflect.TypeOf()
- menggunakan reflect.ValueOf.Kind()


langsung saja ya kita coba

*/

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	var iniString string = "hello"
	var iniInt8 int8 = 10
	var iniSlice []string = []string{"senin", "selasa"}

	// menggunakan %T printf
	fmt.Printf("ini string = %T \r\n", iniString)
	fmt.Printf("ini int8 = %T \r\n", iniInt8)
	fmt.Printf("ini slice = %T \r\n", iniSlice)
	fmt.Println(strings.Repeat("-", 30))

	// penggunakan reflect.TypeOf
	fmt.Println("ini string =", reflect.TypeOf(iniString))
	fmt.Println("ini int8 =", reflect.TypeOf(iniInt8))
	fmt.Println("ini slice =", reflect.TypeOf(iniSlice))
	fmt.Println(strings.Repeat("-", 30))

	// menggunakan reflect.ValueOf.Kind()
	fmt.Println("ini string = ", reflect.ValueOf(iniString).Kind())
	fmt.Println("ini int8 = ", reflect.ValueOf(iniInt8).Kind())
	fmt.Println("ini slice = ", reflect.ValueOf(iniSlice).Kind())

	/*

		dari ketiga cara tersebut, cara ketika berbeda sendiri untuk hasilnya yang
		slice

		kalo type data slice pada refect.ValueOf.Kind() itu akan mengembalikan
		namanya yaitu slice

		akan tetapi 2 yang lain akan mengembalikan type data yang kita deklarasikan
		setelah nama variablenya


	*/

}
