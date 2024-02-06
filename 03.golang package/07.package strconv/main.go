/*


- package strconv

dimateri kali ini kita akan bahas mengenai package strconv
jadi package ini digunakan untuk konversi pada type data string

langsung saja kita coba



*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// mengubah type data string ke int16

	var angka, errParseInt = strconv.ParseInt("100", 10, 16)
	if errParseInt == nil {
		fmt.Println(angka)
	} else {
		fmt.Println(errParseInt.Error())
	}

	/*

		sebelumnya kan untuk konversi type data kita hanya konversi misalnya int8 ke int32
		kita belum coba untuk konversi dari string ke int

		kita gunakan function ParseInt dengan 3 buah parameter
		parameter pertama itu string yang mau kita ubah ke int
		parameter ke 2 itu kita ubah dalam bilangan apa, desimal, oktal, atau hexa
		kalo desimal kalian tuliskan basenya 10

		parameter ke 3 itu mau kita ubah menjadi type data int berapa byte
		misalkan 16 byte jadi yang dikembalikan harusnya type data int16

		function ini mengembalikan 2 buah data, ada data int dan data error
		kenapa ada data error ? karena kan kalo kita memasukkan karakter itu
		akan terjadi error

		untuk function yang lain kalian bisa coba sendiri
		intinya ketika awalnya itu Parse maka function tersebut akan mengubah dari
		type data string ke type data yang diiginkan

	*/

	// mengubah type data int ke string

	var text string = strconv.FormatInt(100, 10)
	fmt.Println(text)

	/*

		jadi kita gunakan function dengan awalan Format
		function dengan awalan Format itu berfungsi untuk mengubah dari type data apapun
		menjadi type data string

		disini dia memiliki 2 buah parameter, parameter pertama itu adalah angkanya
		dan parameter ke dua itu adalah base nya, mau kita ubah ke apa
		hexadesimal atau desimal

	*/

	// menggunakan Atoi dan Itoa

	var text2 string = strconv.Itoa(100)
	fmt.Println(text2)
	var angka2, errAtoi = strconv.Atoi("100")
	if errAtoi == nil {
		fmt.Println(angka2)
	} else {
		fmt.Println(errAtoi.Error())
	}

	/*

		jadi functon Itoa itu digunakan untuk mengubah type data int ke string
		bedanya dengan PerseInt dia tidak perlu memasukkan base dan type data int apa
		jadi lebih singkat

		untuk Atoi mengubah type data string ke int

		jadi seperti itu ya package strconv
		mudah muahan kalian paham

	*/

}
