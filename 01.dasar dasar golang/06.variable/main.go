/*


- variable

sekarang kita akan bahas mengenai variable

variable adalah sebuah tempat yang memiliki nama dan digunakan untuk menyimpan
sebuah value

digolang untuk membuat variable kita harus memberitahu dulu data apa yang mau
kita simpan, data integer,boolean dan lain lain

setelah kita memberitahu type data apa yang mau disimpan
barulah kita simpan type datanya

dan kalo kalian sudah membuat nama variable nya
nama variable tersebut tidak dapat digunakan untuk membuat variable yang baru lagi

untuk membuat variable kita gunakan kata kunci var setelah itu nama variablenya
lalu type datanya dan valuenya
lagsung saja kita coba

*/

package main

import "fmt"

func main() {

	// contoh 1
	var nama string
	nama = "muiz"
	fmt.Println(nama)

	// contoh 2
	var namaku string = "muiz"
	fmt.Println(namaku)

	/*

		untuk mengubah datanya kalian tinggal tulis seperti ini

	*/

	namaku = "hasan"
	fmt.Println(namaku)

	/*

		maka variable namaku akan berubah menjadi hasan
		dan data yang kalian masukkan harus string

		karena kita sudah memberikan type data string kepada variablenya
		kalo kalian tetap memberikan data interger, maka akan error


	*/

	// contoh 3

	var say = "halo"
	fmt.Println(say)

	/*

		kalo kita tidak definisikan type datanya
		maka golang akan secara otomatis mendeteksi type data dari value yang kita
		masukkan itu apa

	*/

	// contoh 4
	umur := 10
	fmt.Println(umur)

	/*

		jadi kalian bisa membuat variable dengan menggunakan titik 2 dan sama dengan
		diikuti dengan datanya

	*/

	// contoh 5 deklarasi multiple variable

	var (
		name string = "muiz"
		age  int8   = 10
	)

	fmt.Println("hello", name, "umur saya", age)
}
