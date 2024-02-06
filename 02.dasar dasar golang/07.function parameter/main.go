/*


- function sebagai parameter

jadi simplenya kita bisa memasukkan function sebagai parameter
ini mirip dengan fitur javascript yaitu callback ya

misalkan saya ingin membuat function cetakNamaFilter

jadi saya ingin mencetak nama diterminal dan melakukan filter
oke langsung saja kita coba


*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func cetakNamaFilter(nama string) string {
	var isUpper bool = false

	var run bool = true
	var i int = 0

	for i < len(nama) && run {
		var s rune = int32(nama[i])
		if unicode.IsUpper(s) {
			isUpper = true
			run = false
		}

		i++
	}

	if isUpper {
		return strings.ToLower(nama)
	}
	return nama

}

func onlyLower(str string) string {
	var isUpper bool = false

	var run bool = true
	var i int = 0

	for i < len(str) && run {
		var s rune = int32(str[i])
		if unicode.IsUpper(s) {
			isUpper = true
			run = false
		}

		i++
	}

	if isUpper {
		return strings.ToLower(str)
	}
	return str
}

type Filter func(string) string

func cetakNamaFilter2(nama string, filter Filter) {
	var cetak string = filter(nama)
	fmt.Println(cetak)
}

func main() {

	var result string = cetakNamaFilter("muiz")
	fmt.Println(result)

	/*

		jadi kita punya function cetak nama dimana function tersebut memiliki filter
		kalo misalkan ada karaketer yang memiliki huruf besar maka ubah string tersebut
		menjadi lower case dan kembalikan

		kalau tidak langsung return saja

		akan tetapi disini ada masalah nih

		yang pertama function tersebut memiliki 2 tugas yaitu mencetak nama dan melakukan filter
		jadi ketika membuat function sebaiknya kita melakukan satu tugas tertentu saja

		jadi satu function memiliki satu tugas

		yang kedua, function ini memiliki nama filter dimana filter ini tidak spesifik
		dan bisa filter apapun

		masalah pertama dapat kita atasi dengan membuat 2 function
		function pertama cetak nama dan function kedua filter

		masalah yang kedua kita bisa atasi dengan menambahkan function filter yang telah
		kita buat sebagai parameter

		langsung saja kita coba
	*/

	cetakNamaFilter2("muIz", onlyLower)

	/*

		nah jadi bisa seperti ini,
		jadi functionnya terpisah dan mengerjakan satu tugas masing masing

		jadi bacanya gini, function onlyLower sama masukkan kedalam
		parameter filter

		lalu kemudian saya jalankan function filter dan saya kirimkan argument
		nama

		lalu function filter akan mengembalikan
		string dan mencetak kelayar

		jadi seperti itu ya, kalian juga bisa menambahkan type deklaration
		pada parameternya

	*/

}
