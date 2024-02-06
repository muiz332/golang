/*

- closure


sekarang kita akan bahas mengenai closure

apa itu closure ?
closure adalah kemampuan sebuah function untuk mengakses data pada scop yang diatasnya

akan tetapi kalian harus berhati hati bisa jadi secara tidak sengaja
kalian bisa mengubah value dari sebuah variable lain dan itu bisa kalian ubah
didalam sebuah function

dari pada bingung langsung saja kita coba



*/

package main

import "fmt"

func main() {

	var nama string = "hasan"

	var ubahNama = func() {
		nama = "husain"
	}

	ubahNama()
	fmt.Println(nama)

	/*

		kalo kita jalan, makan variable nama akan berubah
		karena kita bisa ubah data didalam sebuah function

		kenapa bisa seperti ini ?
		ketika kian membuat sebuah function

		dan ketika ada proses assigment didalam function tersebut
		maka golang akan mencari variable tersebut

		ada kah deklarasi variable nama didalam function ubahNama ? tidak ada
		maka dia akan mencari pada scrop atau lingkup diluar function pada code
		diatas function tersebut

		ketika mencari keluar, ternya ada nih deklarasi variable nama
		maka nama kita ubah menjadi husain

		kalo kita tampilkan kelayar, maka nama sudah berubah valusenya menjadi husan
		inilah yang dimanakan dengan closure

		kemampuan sebuah function bisa mengakses scop yang diatasnya
		lalu bagaimana kalo kita tidak ingin mengubah variable nama

		kita tinggal buat deklarasinya didalam sebuah function ubah nama

	*/

	var nama2 string = "hasan"

	var ubahNama2 = func() {
		var nama2 string = "hasan"
		nama2 = "husain"
		fmt.Println("ini dari function", nama2)
	}

	ubahNama2()

	fmt.Println(nama2)
	/*

		maka kalo kita jalankan variable nama2 yang ada diluar function tidak
		akan berubah

		jadi seperti itu ya, gunakan fitur ini dengan sangat hati hati
		dan juga kalian bisa mengakses function untuk dipanggil kedalam sebuah function

		akan tetapi kalo sebuah variable didalam block function
		block maksutnya kurung kurawal buka dan tutup

		maka variable tersebut tidak akan bisa diakses diluar function tersebut
		jadi seperti itu ya closure

		mudah mudahan kalian paham

	*/

}
