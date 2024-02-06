/*


- operasi logika

operasi logika adalah sebuah operasi yang mambandingkan dua atau lebih kondisi
untuk kondisi itu seperti materi sebelumnya ya


macam macam operasi logika


&& and
|| or
! not


langsung saja kita coba


*/

package main

import (
	"fmt"
)

func main() {

	// operator logika AND

	var nilai int8 = 80
	var absen int8 = 10
	var lulus bool = nilai >= 70 && absen >= 8
	fmt.Println(lulus)

	/*

		jadi bacanya gini, saya memiliki variable nilai dengan value 80
		dan variable absen dengan nilai 10

		saya ingin check apakah dia lulus atau tidak?
		apakah nilainya itu lebih besar 70 true
		dan
		apakah absennya itu lebih besar 8 true

		maka jika keduanya true maka hasilnya true
		itu operator logika and

	*/

	// operator logika OR

	var angka int8 = 20
	var hasil bool = angka <= 10 || angka >= 20
	fmt.Println(hasil)

	/*

		apakah angka itu lebih kecil sama dengan 20? false
		atau
		apakah angka itu lebih besar sama dengan 20? true

		jadi ketika ada salah satu kondisi yang bernilai true
		maka hasilnya akan true, jika semua false maka hasilnya false


	*/

	// operator logika NOT

	var status bool = true
	fmt.Println(!status)

	/*

		jadi kalo kondisinya true dibalik jadi false
		kalo false dibalik jadi true

	*/
}
