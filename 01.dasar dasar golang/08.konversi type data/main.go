/*

- konversi type data

kadang kita juga perlu melakukan konversi type data
misalkan dari int16 ke int8, atau dari float ke int32

untuk melakukan konversi type data digolang cukup sederhana
kalian bisa lakukan seperti ini


*/

package main

import "fmt"

func main() {
	var angka1 int16 = 130
	var angka2 int32 = int32(angka1)
	var angka3 int8 = int8(angka2)

	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(angka3)

	/*

		jadi bacanya, dari variable angka1, saya ingin ubah menjadi int32,
		cukup kalian tuliskan function int32() dan kalian tambahkan variablenya
		didalam kurung

		maka yang dikembalikan kedalam angka2 adalah int16

		akan tetapi kalian harus berhati hati kalo angka yang dimasukkan
		melebihi kapasitas dari type datanya

		misalkan variable angka3, itukan type datanya int8, sedangkan int8 maximal
		valunya adalah 127, sedangkan angka yang kita masukkan adalah 130

		kelebihan 3 nilai tuh, kalo kelebihan maka dia akan mengulang dari nilai
		paling kecilnya

		nilai paling kecil dari int8 adalah -128
		jadi -128 dikurangi 3 adalah -126

		maka nilainya adalah -126

		kasus seperti ini disebut dengan number overflow,
		dimana kita melakukan konversi type data melebihi kapasitas dari type data
		yang ingin kita konversi


		ada contoh lagi seperti ini, misalkan kita ingin mengambil karakter dari
		string, kita bisa menggunakan kurung kotak dan didalamnya terdapat posisi
		karakter yang mau kita ambil, dan dimulai dari 0,

		itu disebut index, akan tetapi yang muncul adalah angka bukan karakter yang
		kita inginkan, kita bisa konversi angka tersebut menjadi string
	*/

	var nama string = "muiz"
	var unicode byte = nama[0] // byte itu unit8
	var karakter string = string(unicode)

	fmt.Println(nama[0])  // 109
	fmt.Println(karakter) // m

	/*

		kenapa muncul angka? tidak langsung muncul karakter saja seperti dijavascript?
		karena golang akan memunculkan yang namanya unicode dari sebuah karakter

		unicode adalah standart pengkodean karakter yang digunakan untuk menrepresentasikan
		karaketer dan simbol dalam berbagai bahasa dan penulisan diseluruh dunia

		jadi setiap karakter atau simbol itu memiliki unicode masing masing,
		golang memunculkan unicode tersebut dari tiap tiap karakter

	*/

}
