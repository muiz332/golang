/*

-  defer

apa itu defer?
defer adalah sebuah funciton yang dijalankan setelah sebuah function selesai
dieksekusi meskipun terjadi error

kadang kadang kita butuh mengeksekusi sebuah function meskipun terjadi error

langsung kita coba

*/

package main

import "fmt"

func endFunc() {
	fmt.Println("program selesai dijalankan")
}

func devided(num int, devisor int) {
	defer endFunc()

	fmt.Println(num / devisor)
}

func main() {

	devided(10, 2) // 5

	/*

		selalu inget untuk menuliskan defer functionnya dipaling atas
		sendiri

		kalo saya jalankan maka akan tampil hasilnya 5 dan akan muncul
		program selesai dijalankan

		akan tetapi kita coba kalo misalkan terjadi error
		apakah function endFunc yang kita jadikan sebagai
		defer function akan tetap dijalankan meskipun error ?

		kita coba misalkan 10 / 0, jadi kalo dibagi 0 itu digolang
		langsung error ya

	*/

	devided(10, 0)
	fmt.Println("berhenti")

	/*

		kalo kita jalankan muncul seperti ini ?
		5
		program selesai dijalankan
		program selesai dijalankan

		jadi setelah pemanggilkan function pertama setesai
		maka akan langsung muncul, program selesai dijalankan
		dan akan tampil error seperti ini

		panic: runtime error: integer divide by zero

		dan tulisan berhenti tidak akan tampil kelayar

		jadi kesimpulannya, defer function itu bisa kita gunakan untuk tetap menjalankan
		sebuah function setelah sebuah function selesai dieksekusi meskipun terjadi error

		akan tetapi setelah terjadi error maka program akan terhenti
		dan tulisan berhenti tidak akan ditampilkan dilayar

		sekarang gimana kalo saya ingin menggantti tulisan errornya
		misalkan tulisan errornya menjadi

		angka tidak bisa dibagi dengan 0

		nah kita bisa gunakan sebuah funciton yang namanya panic

	*/

}
