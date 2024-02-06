/*


- slice function

nah kali ini kita akan coba beberapa function yang sering kita gunakan
didalam slice

len() 							mengecheck total data dari slice
cap() 							mengecheck capacity dari slice
append(slice,data) 				menambah data didalam array
make([]string,length,capacity)  membuat slice baru
copy(destination, source)  		mengcopy slice dari source ke destination

langsung saja kita coba satu persatu ya

*/

package main

import "fmt"

func main() {

	var hari = [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	var slice1 []string = hari[3:5]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println()

	// append

	slice1 = append(slice1, "tambah")
	fmt.Println(slice1)
	fmt.Println(hari)
	fmt.Println()

	/*

		kalo kita jalankan maka hasilnya akan seperti ini
		[kamis jumat tambah]
		[senin selasa rabu kamis jumat tambah minggu]

		slicenya bertambah, akan tetapi kenapa data sabtu diarray hari
		itu tidak ada, kenapa malah diganti dengan data tambah?

		jadi digolang memang seperti itu
		ketika kita menambah menambah slicenya menggunakan append, maka slicenya
		memang akan bertambah

		akan tetapi arraynya akan berdampak, dalam kasus ini
		pointer dari slicenya yaitu 3 sampai index ke 4


		ketika kita tambah datanya otomatis slicenya akan menjadi
		index k3 sampai 5, karena slice itu mereference array yang
		kita potong

		maka didalam array index ke 5 yang tadinya datanya sabtu berubah
		menjadi data yang baru saja kita tambahkan yaitu data tambah

		jadi slicenya tetap mengacu ke array yang dipotong


		nah sekarang pertanyaannya, kalo misalkan kita menambah data di
		slicenya melebihi kapasitas dari slice itu sendiri

		apa yang akan terjadi?
		kita coba, kita tambahkan 2 data lagi, karena capacitynya 4

		jadi untuk menambahkan data didalam slicenya itu kalian
		tuliskan variable slicenya kemudian function append

		jadi secara tidak langsung variable slice tersebut direplace
		dengan data yang sudah kita tambahkan menggunakan append

		oke, kita coba
	*/

	slice1 = append(slice1, "kurang", "bagi")
	fmt.Println(slice1)
	fmt.Println(hari)
	fmt.Println()

	/*

		kalo kita jalankan maka hasilnya seperti ini

		[kamis jumat tambah kurang bagi]
		[senin selasa rabu kamis jumat tambah minggu]

		nah sekarang minggunya tidak berubah, padahal datanya kita tambah
		kenapa seperti ini ?


		digolang ketika kita menambahkan data didalam slice, maka golang akan
		mengecheck

		apakah data yang kita tambahkan melebihi kapasitasnya, kalo tidak dia
		akan menambah datanya dan juga mereplace array yang menjadi acuannya

		kalo iya, maka dia akan membuat array baru dibelakang layar
		dan ketika dia sudah membuat array baru dibelakang layar

		maka acuannya berpindah dari yang array hari, sekarang menjadi array yang
		dibelakang layar tersebut

		lalu apa buktinya?
		kita coba ubah data didalam array hari,
		kalo slicenya berubah maka dia tetap mengacu kepada array hari

		kalo tidak berarti dia mengacu kepada array yang sudah dibuat
		dibelakang layar

		kita coba

	*/

	hari[3] = "kamiskuy"
	fmt.Println(hari)
	fmt.Println(slice1)

	/*
		kalo kita jalankan hasilnya seperti ini

		[senin selasa rabu kamiskuy jumat tambah minggu]
		[kamis jumat tambah kurang bagi]

		array harinya berubah menjadi kamiskuy, akan tetapi
		slicenya tidak berubah, tetap menjadi kamis

		ini berarti acuannya tidak lagi diarray hari melainkan sudah berpindah
		diarray yang telah dibuat secara otomatis dibelakang layar ketika
		kita menambahkan datanya didalam slice melebihi capacitynya


		jadi seperti itu tentang function append


		kita lanjutkan penggunaan function make
		untuk membuat slice
	*/

	slice2 := make([]string, 2, 10)
	slice2[0] = "muiz"
	slice2[1] = "hasan"

	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println()

	/*

		nah untuk membuat slice yang bisa menentukan length dan capacitynya
		kita bisa menggunakan function make seperti itu ya

		kalo kalian membuat seperti ini, dibelakang layar maka akan dibuatkan
		array dengan panjang 10, ini akan lebih aman karena arraynya diumpetin
		dibelakang layar


		selanjutnya kita coba copy slice
	*/

	slice3 := make([]string, len(slice2), cap(slice2))

	copy(slice3, slice2)
	fmt.Println(slice3)

	/*

		jadi sebelum kita copy slicenya kita buat dulu ya slicenya
		misalkan slice3, dimana type datanya slice string, len dan capacitynya
		harus sesuai dengan slice yang ingin kita copy

		kalo lebih sedikit, nanti slice2nya tidak akan dicopy semuanya
		jadi seperti terpotong

		jadi misalkan lenghtnya 1 maka yang dicopy hanyalah value muiz saja
		dari slice2

		jadi bacanya
		saya copy slice2 ke sclie3



	*/

}
