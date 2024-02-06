/*


- function return value


sebuah function biasanya seperti ini


inputan (argument)

proses

hasil


yang pertama kali dilakukan function adalah saat dia dipanggil jika ada parameternya
maka data dari parameter tersebut akan diproses didalam function

setelah diproses, maka function akan mengembalikan hasilnya
hasil itulah yang kita sebut dengan return value

bagaimana function bisa mengembalikan nilai kita bisa menggunakan keyword
return dan diikutin dengan value yang kita kembalkan

misalkan saya ingin membuat function yang fungsinya untuk
menjumlahkan dua buah nilai

nilai a dan b setelah itu function tersebut akan mengembalikan hasil
dari penjumlahan

langsung saja kita coba

*/

package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func returnMultipleValue() (string, string) {

	return "muiz", "hasan"
}

func returnMultipleValueDifferent() (string, int8) {
	return "muiz", 10
}

func main() {

	fmt.Println(sum(1, 2)) // 3

	/*

		jadi setelah kita tambah a dan b maka kita return valuenya
		dan jangan lupa kalian harus memberikan type data untuk return valuenya

		setelah itu kalian bisa print untuk memunculkan hasilnya

		atau kalian bisa masukkan data tersebut kedalam sebuah variable


	*/

	var hasil int = sum(2, 2)
	fmt.Println(hasil) // 4

	/*

		jadi bisa seperti itu,
		itulah gunanya return value didalam function

		return selain mengembalikan data, dia juga bisa menghentikan program didalam function
		kalo kalian menuliskan code dibawah return

		maka code tersebut tidak akan pernah dijalankan


		selanjutnya digolang juga bisa mereturn lebih dari satu data
		dengan menambahkan kurung buka dan tutup serta type data lagi setelah kurung buka dan tutupnya

		jika lebih dari satu kita bisa memisahkannya menggunakan koma
		langsung saja kita coba

	*/

	var nama1, nama2 string = returnMultipleValue()
	fmt.Println(nama1, nama2)

	/*

		jadi seperti itu ya, kalo misalkan type datanya sama
		maka sebelum tanda sama dengan kalian bisa tambahkan type datanya
		yaitu string

		akan tetapi kalo type datanya berbeda kalian bisa deklarasikan
		terlebih dahulu variablenya, barulah kita assigment

	*/

	var myName string
	var age int8

	myName, age = returnMultipleValueDifferent()
	fmt.Println(myName, age)

	/*

		jadi seperti itu ya kalo type data dari return valuenya berbeda
		kalian bisa deklarasikan type datanya

		atau kalian bisa membuat variable dengan operator :=

		dan kalo kalian tidak butuh salah satu dari variable tersebut
		kalian bisa ubah variable tersebut menjadi tanda underscore _

	*/

	myName1, _ := returnMultipleValue()
	fmt.Println(myName1)

}
