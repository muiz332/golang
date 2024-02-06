/*

- function value



function adalah first class citizen digolang
function juga merupakan type data dan bisa disimpan didalam sebuah variable


langung sama kita coba

*/

package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

var goodMorning func(nama string) = func(nama string) {
	fmt.Println("good morning", nama)
}

func main() {

	var iniSayHello func() = sayHello
	iniSayHello()

	/*

		jadi seperti itu ya, kita bisa memasukkan sebuah function
		kedalam sebuah variable

		dan ingat kalian jangan menuliskan kurung buka dan tutupnya saat melakukan
		assigment kedalam variable, karena itu akan langsung menjalankan
		functionnya

		atau kalian juga bisa membuat function tanpa nama dan langsung dimasukkan kedalam
		sebuah variable
	*/

	goodMorning("muiz")

	/*

		jadi digolang function itu merupakan type data, yang artinya kita bisa
		memasukkan kedalam sebuah variable

		kalo seperti ini, sebuah function juga dapat kita masuakkan ke dalam parameter
		menjadi sebuah value

	*/
}
