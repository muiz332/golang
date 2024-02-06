/*

- if


saat membuat program, ada kalianya kita ingin mengatur control flow
dari sebuah program

misalkan ketika ada sebuah program menemui angka dibawah 4
lakukan hal ini, atau ketika diatas 4 lakukan hal ini

jadi kita bisa mengatur jalannya program kita, itulah yang dinamakan
percabangan atau kalo dibahasa pemrograman kita menyebutnya if, atau
pengkondisian

pengkondisian adalah kita bisa mengeksekusi code program kita
ketika suatu kondisi terpenuhi, jika tidak terpenuhi maka kode
program tidak akan dijalankan

untuk membuat pengkondisian kita tulis seperti ini


if kondisi {
	aksi
}

jadi kalian bisa memasukkan kondisi, untuk kondisinya itu
seperti kondisi yang sudah pernah kita bahas ya
dimateri operasi perbandingan dan operasi logika

untuk aksi, setelah kondisi tersebut bernilai true
maka aksi akan dijalankan

aksi ini, adalah code program apa yang akan kalian jalankan
misalkan memunculkan hello world kelayar

untuk lebih jelasnya langsung saja kita coba

misalkan saya ingin mengecheck apakah
sebuah angka itu bilangan ganjil atau genap

kalo bilangan genap itu ketika saya modulus dengan 2
hasilnya akan 0

kita coba


*/

package main

import "fmt"

func main() {

	var angka int8 = 10

	if angka%2 == 0 {
		fmt.Println(angka, "adalah angka genap")
	}
	/*

		kalo kita jalankan maka akan memunculkan print kelayar
		10 aadalah angka genap

		karena kondisinya bernilai true, maka aksinya dijalankan
		kalau false, maka aksinya tidak akan pernah dijalankan

		akan tetapi saya juga ingin menambahkan kondisi
		ketika saya modulus 2 sisa baginya 1, maka munculkan
		angka, adalah angka ganjil

		maka dari itu kita membutuhkan kondisi lagi,
		apakah kita akan menambahkan if lagi?

		sebenarnya bisa saja, tetapi ini sangat tidak direkomendasikan
		karena yang kita check variable yang sama

		untuk menambahkan kondisi kita gunakan keyword else if

	*/

	var angka2 int8 = 11

	if angka2%2 == 0 {
		fmt.Println(angka2, "adalah angka genap")
	} else if angka2%2 == 1 {
		fmt.Println(angka2, "adalah bilangan ganjil")
	}

	/*

		jadi bacanya gini,
		apakah angka saya modulus 2 sama dengan 0 ? false
		maka kita skip

		lalu ada kondisi lagi nih, apakah angka saya modulus dengan 2
		sama dengan 1  ? true, maka dia akan menjalankan aksi yang ada
		didalam kurung else ifnya

		kalian juga bisa menambahkan else if lagi jika memang dibutuhkan


		selanjutnya ada keyword else, else itu digunakan ketika ingin menjalankan
		sesuatu selain kondisi yang diatasnya,


		jadi kalo misalkan kondisi yang sudah kita buat tidak terpenuhi atau
		bernilai false semua, maka else akan dijalankan

		kita coba

	*/

	var nama string = "husain"

	if nama == "muiz" {
		fmt.Println("hello", nama)
	} else if nama == "hasan" {
		fmt.Println("hai", nama)
	} else {
		fmt.Println("siapa anda ?")
	}

	/*

		jadi apakah nama sama dengan muiz? false
		apakah nama sama dengan hasan ? false

		selain itu saya ingin menjalankan apapun yang ada
		didalam elsenya


		selanjutnya ada if short statement
	*/

	if umur := 10; umur <= 10 {
		fmt.Println("masih muda")
	} else if umur > 10 && umur < 20 {
		fmt.Println("remaja")
	} else {
		fmt.Println("dewasa")
	}

	/*

		jadi kita bisa membuat variable didalam ifnya secara langsung
		dan kita langsung mengecheck dengan kondisnya



		jadi seperti itu ya penggunaan dari pengkondisian ini
		mudah mudahan kalian paham

	*/
}
