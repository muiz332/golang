/*

- type deklaration



type deklaration itu adalah aliases dari type data yang sudah ada
jadi kita bisa membuat aliases dari type data yang sudah ada

contohnya

byte, byte itu adalah aliases dari type data int8
rune, rune itu aliases dari int

kita juga bisa buat seperti itu
langsung saja ya kita coba

*/

package main

import "fmt"

type NoKTP string

func main() {

	var myKtp NoKTP = "123454656"
	fmt.Println(myKtp)

	/*

		jadi kalian bisa tuliskan type deklaration ini didalam atau
		diluar function main

	*/

}
