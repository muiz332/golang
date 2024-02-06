/*

- recover

recover adalah sebuah funciton yang bisa kita gunakan untuk menangkap data
panic

dengan recover proses panic akan terhenti, sehingga program akan tetap
dijalankan

langsung saja kita coba



*/

package main

import "fmt"

func endFunc() {

	var errMessage = recover()
	if errMessage != nil {
		fmt.Println(errMessage)
		return
	}

	fmt.Println("function telah selesai dijalankan")
}

func devidedBy(num int, devisor int) {
	defer endFunc()

	if devisor == 0 {
		panic(fmt.Sprintf("%d tidak bisa dibagi dengan 0", num))
	}
}

func main() {

	devidedBy(10, 0)
	fmt.Println("program selesai dijalankan")

	/*

		untuk menangkap pesan error yang telah dituliskan difunction panic
		maka kita harus tangkap pada defer function

		dimana didalam function endFunc kita memanggil function recover
		dan function tersebut mengembalikan type data any

		nanti kita akan type data any, akan tetapi intinya yang dikembalikan
		adalah type data string yaitu messagenya

		dan kalo tidak terjadi error yaitu nil, untuk type data nil
		nanti kita akan bahas

		akan tetapi ketika error function recoveri mengembalikan pesan errornya
		dan ketika tidak terjadi error maka akan mengembalikan nil

		jadi kita bisa check tuh menggunakan if, apakah terjadi error ?
		kalo terjadi error maka tampilkan ke layar pesan errornya

		dan ingat return untuk menghentikan program nya
		jadi itu ya mengenai recover

		dan tuliskan program selesai dijalankan akan tampil kelayar
		artinya program tidak akan berhenti

		mudah mudahan kalian paham

	*/
}
