/*

- String

type data string adalah sebuah type data yang memiliki lebih dari satu karakter
dimana karakter tersebut dapat berupa huruf, angka, simbol, tanda baca, dan karaker karakter
yang lainnya

type data string direpresentasikan dengan kata kunci string
digolang string itu menggunakan tanda kutip 2


*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello ini adalah string")

	/*

		nah selanjutkan jika kalian ingin tahu jumlah karakter
		yang ada didalam string, kalian gunakan function
		namanya len()

		jadi kalian bisa tulis seperti ini

	*/

	fmt.Println(len("hello")) // 5

	/*

		maka akan muncul hasilnya 5
		dan kalo kalian ingin mengambil karaketer didalam string tersebut
		maka kalian tuliskan seperti ini

	*/

	fmt.Println("hello"[0]) // 104

	/*

		kalian gunakan kurung kotak kemudian ditengah tengahnya kalian tuliskan
		posisi dari string tersebut, akan tetapi dimulai dari angka 0 ya

		akan tetapi munculnya adalah angka 104
		kita harus ubah angka tersebut menjadi karaketer

		caranya kalian tuliskan seperti ini

	*/

	fmt.Println(string("hello"[0])) // h

	/*

		jadi kalian tinggal ubah saja angka tersebut menjadi string
		dengan menggunakan function string()



	*/
}
