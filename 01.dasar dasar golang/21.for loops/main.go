/*

- for loops



selain if statement, untuk mengatur control flow, atau alur dari program kita,
kita bisa menggunakan for loops


for loop digunakan untuk mengulang code program kita ketika kondisinya
bernilai true

ketika false, maka dia akan berhenti dari pengulangannya

bisanya didalam bahasa pemrograman, pengulangan ada tiga jenis
ada forloop, while, dan do while

digolang hanya ada 1, yaitu for loops, akan tetapi dia bisa melakukan apa
yang while lalukan

untuk membuatnya kalian bisa tulis seperti ini


nilai awal

for kondisi {
	aksi

	increment/decrement
}

itu adalah for dalam bentuk seperti while

untuk for biasa seperti ini

for nilai awal; kondisi; increment/decrement{
	aksi
}

biasanya for dalam bentuk while itu digunakan ketika
kita memiliki kondisi yang complex

kalo for biasa biasanya kondisinya hanya dalam bentuk operator perbandingan
saja, misalkan lebih besar dari, lebih kecil dari dan lain lain


*/

package main

import "fmt"

func main() {

	// for dalam bentuk while

	var angka1 int32 = 0

	for angka1 < 10 {
		fmt.Println("hello world")

		angka1++
	}
	fmt.Println("looping selesai")
	fmt.Println()

	/*

		jadi misalkan saya ingin mengulang kata hello world sebanyak 10 kali
		maka kalian bisa tuliskan nilai awal untuk awal mulainya

		dan kondisinya, kalo misalkan angka lebih kecil 10, karena saya ingin mengulang
		sampai 10 saja,  maka jalankan aksinya

		untuk membuat kondisinya bernilai false, variable angka1 harus bernilai lebih besar sama
		dengan 10

		biar angkanya naik setiap looping, maka kita tambahkan increment yaitu penambahan nilai awal

		angka1++, jadi ini sama dengan angka1 = angka1 + 1
		jadi tiap looping angka1 akan ditambah dengan 1

		jika sudah menuju angka 9, maka dia dicheck
		apakah 9 itu lebih kecil 10 ? true

		maka dia masuk looping dan menjalankan aksinya, dan menjalankan increment
		angka1++

		jadi angka1 sekarang bernilai 10, masuk kedalam kondisi lagi
		dicheck apakah angka1  itu lebih kecil 10 ? false

		karena angka1 itu bernilai 10, kalo false dia akan keluar dari loopingnya dan
		melanjutkan program yang ada dibawahnya

		jadi seperti itu ya,

		kalian juga bisa menambahkan pengkondisian didalam loopingnya,
		misalkan saya ingin mengecheck

		kalo angka1 itu bilangan genap, maka saya akan tampilkan hello world
		kita coba

	*/

	var angka2 int32 = 0

	for angka2 < 10 {
		if angka2%2 == 0 {
			fmt.Println("hello world")
		}

		angka2++
	}
	fmt.Println("looping yang didalamnya ada if selesai")
	fmt.Println()

	/*

		maka muncul tuh hasilnya akan ada 5 kata hello world
		akan tetapi kalian harus hati hati

		kalo kalian lupa menuliskan increment atau decrementnya
		maka loopingnya akan berjalan terus dan tidak akan berhenti

		jadi kalo kondisinya tetap true dia akan berjalan terus
		jadi hati hati ya


		selanjutnya kita kan coba menggunakan for biasa


	*/

	for angka1 := 10; angka1 > 0; angka1-- {
		fmt.Println("ini looping ke", angka1)
	}
	fmt.Println("loop selesai")

	/*

		sama saja, akan tetapi nilai awal dan increment atau decrementnya saya gabung menjadi
		satu baris,

		jadi gini bacanya
		masuk ke for

		angka saya isi dengan nilai 10, dia akan menjalankan nilai
		awal hanya satu kali

		lalu dicheck apakah angka lebih besar 0 ? true
		lalu dia akan menjalankan aksinya yaitu munculkan tulisan kelayar

		lalu dia akan menjalankan decrement yaitu pengurangan nilai awal
		karena saya ingin mengulang dari 10 sampai 0

		maka angka1 sekarang menjadi 9
		terus seperti itu sampai angka1 bernilai 0

		setelah itu dicheck, apakah 0 itu lebih besar 0 ? false
		maka loopingnya akan berhenti, dan akan melanjutkan program
		selanjutnya

		yaitu memunculkan kata loop selesai diterminal

		dan kalian juga bisa menambahkan didalam for itu for lagi
		jadi tidak ada batasnya ya

		mau didalamnya diberikan for lagi yang didalamnya dikasih for lagi
		bebas saja tergantung kebutuhan ya

	*/

}
