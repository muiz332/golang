/*

- operasi aritmatika

operasi aritmatika atau matematika digolang itu sama seperti bahasa
pemrograman yang lain

macam macam operasi aritmatika

-  pengurangan
+  penjumlahan
/  pembagian
*  perkalian
%  modulus atau sisa bagi


langsung saja kita coba


*/

package main

import "fmt"

func main() {

	var angka1 int8 = 8
	var angka2 int8 = 2
	var hasil int8 = angka1 / angka2

	fmt.Println(hasil)

	/*

		jadi dari variable angka1 dibagi dengan variable angka2
		maka hasilnya akan dimasukkan kedalam variable hasil

		dan akan dicetak kelayar menggunakan print


		nah selanjutnya ada operasi yang lain namanya operasi assigment
		seperti

		+=
		-=
		/=
		*=
		%=

		langsung saja kita coba

	*/

	var angka3 int8 = 10
	angka3 += 10 // sama saja dengan angka3 = angka3 + 10
	fmt.Println(angka3)

	/*

		jadi kita bisa langsung menambahkan angka dari variable yang sudah kita buat
		menggunakan operator +=

		jadi dia akan melakukan penambahkan terhadap variablenya sendiri kemudian hasilnya
		akan dimasukkan kedalam variable tersebut


		selanjutnya ada unary operator

		++      sama saja dengan   a = a + 1
		--      sama saja dengan   a = a - 1
		-                          negatif
		+                          positif
		!                          boolean kebalikan
	*/

	angka3++
	fmt.Println(angka3)

	fmt.Println(-angka3)
	fmt.Println(+angka3)
	fmt.Println(!true)

}
