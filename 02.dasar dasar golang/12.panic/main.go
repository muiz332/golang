/*


- panic

apa itu panic?
panic adalah sebuah function yang dapat kita gunakan untuk menuliskan error secara
spesifik

jadi kita bisa mengcustom tulisan errornya, atau kita bisa membuat error itu sendiri
kita coba ya


*/

package main

import "fmt"

func endFunc() {
	fmt.Println("function selesai dijalankan")
}

func devidedBy(num int, devisor int) {
	defer endFunc()

	if devisor == 0 {
		panic("angka tidak bisa dibagi dengan 0")
	}

	fmt.Println(num / devisor)
}

func main() {

	devidedBy(10, 0)
	fmt.Println("program selesai")

	/*

		kalo kita jalankan maka akan defer function akan tetap dijalankan
		dan errornya akan muncul seperti ini

		panic: angka tidak bisa dibagi dengan 0

		jadi akan muncuk panic: angka tidak bisa dibagi dengan 0
		sesuai dengan apa yang kita tuliskan di function panic ya

		jadi itulah panic, kita bisa membuat error dengan menggunakan
		function panic

		jadi kalo didalam javascript, panic ini seperti dengan
		throw new Error("angka  tidak bisa dibagi dengan 0")

		jadi seperti itu ya panic, ketika panic maka program akan terhenti
		akan tetapi ketika terjadi error saya ingin

		program tetap dijalankan, dan tulisan error saya ingin tampilkan kelayar
		jadi ada 2 poin nih

		bagaimana kita membuat error tanpa menghentikan program
		dan bagaimana kita bisa menangkap pesan error dari panic
		dan menampilkannya kelayar ?


	*/
}
