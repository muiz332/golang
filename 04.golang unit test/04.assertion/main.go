/*

- assertion


melakukan pengecheckan unit test secara manual menggunakan if else itu sangat ribet
apa lagi jika result datanya banyak

misalkan kalian memiliki struct yang  propertinya ada 4
jadi kalian harus bikin kondisi if else sebanyak 4 kali

oleh karena itu sangat disarankan untuk menggunakan assertion ini
golang tidak memiliki package assertion sehingga kita butuh menginstall package pihak
ketiga untuk melakukan assertion ini

nama packagenya adalah testify

untuk menginstallnya kalian cukup tuliskan

go get github.com/stretchr/testify


langsung saja kita coba




*/

package main

import (
	"fmt"

	say "github.com/Muizzuddin-github/sayhello"
)

func main() {

	var nama = say.SayHello()
	fmt.Println(nama)

}
