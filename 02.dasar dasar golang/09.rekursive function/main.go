/*


- rekursive function

apa itu rekursive funciton ?
rekursive function adalah function yang memanggil dirinya sendiri

nah sebenarnya didalam function kita bisa memanggil function yang lain
akan tetapi kalo misalkan kita memanggil function diri sendiri

maka itu yang dikatakan rekursive function
contohnya kita bisa menyelesaikan kasus faktorial menggunakan rekursive function


akan tetapi kita coba menggunakan for loop dulu
barulah kita coba menggunakan rekursive function

*/

package main

import "fmt"

func faktorialLoop(angka int) int {
	var result = 1
	for i := 2; i <= angka; i++ {
		result *= i
	}

	return result
}

func rekursive(angka int) int {

	if angka == 1 {
		return 1
	}

	return angka * rekursive(angka-1)
}

func main() {

	// rekursive menggunakan for loop

	var result int = faktorialLoop(5)
	fmt.Println(result) // 120

	/*

		nah sekarang kita coba yang rekursive function
		saat sebuah function memanggil dirinya sendiri

		maka dia memanggil function dirinya sendiri tanpa batas
		kita harus menghentikan pemanggilan dirinya sendiri

		ketika kondisi tertentu telah tercapai,
		ini yang disebut dengan base case

		jadi base case ini adalah sebuah kondisi dimana kita
		bisa menghentikan rekursive function

		kita coba ya

	*/

	var hasil int = rekursive(5)
	fmt.Println(hasil) // 120

	/*

		jadi bacanya gini

		ketika funciton rekursive dijalankan
		dia masuk kedalam kondisinya

		apakah angka itu sama dengan 1? false

		maka dia akan memanggil function rekursive dimana parameter angkanya
		dikurangi 1

		jadi terus memanggil function rekursive dan mengurangi parameter angka
		sampai ke angka 1

		ketika sampai ke angka 1
		maka dicheck tuh dikondisinya

		apakah parameter angka itu sama dengan 1 ? true
		kalo true kita return 1

		jadi return itu menghentikan funcitionnya ya

		setelah direturn dia akan masuk ke pemanggilan function berikutnya
		yaitu angka sama dengan 2

		maka akan dikalikan 2 * 1 ? hasilnya 2
		terus seperti itu sampai angka ke 5

		2 * 1 = 2
		3 * 2 = 6
		4 * 6 = 24
		5 * 24 = 120

		jadi itulah rekursive function
		mudah mudahan kalian paham


	*/
}
