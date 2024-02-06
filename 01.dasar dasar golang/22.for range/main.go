/*


- for range


biasanya kita bisa menggunakan for untuk menelusuri array, slice, ataupun map
kecuali map

kalau mapnya itu memiliki key angka yang tidak berurutan
atau memiliki keynya type data selain angka

maka hanya bisa menggunakan for range

kita coba



*/

package main

import "fmt"

func main() {

	// for biasa

	var hari [7]string = [7]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	for i := 0; i < len(hari); i++ {
		fmt.Println(hari[i])
	}
	fmt.Println()

	/*

		jadi kita menggunakan function len untuk mengetahui seberapa banyak panjang arraynya
		karena kita ingin menelusuri array mulai dari index ke 0 sampai index terakhir array
		yaitu index ke 7

		kalo kita jalankan maka akan muncul ya
		senin
		selasa
		rabu
		kamis
		jumat
		sabtu
		minggu

		jadi seperti itu
	*/

	//  foor range

	for i, val := range hari {
		fmt.Println(val, i)
	}
	fmt.Println()

	/*

		jadi seperti itu ya, ada dua buah variable setelah for
		ada i sebagai indexnya, dan ada val sebagai valuenya atau datanya

		jadi kita bisa menampilkan nya dilayar

		akan tetapi kalo misalkan kita tidak menggunakan variable i maka
		golang akan mendeteksi error

		lalu bagaimana kalo misalkan saya tidak butuh variable i ?

		kalian bisa ganti variable i dengan menggunakan tanda underscore _
		agar kita bisa mengabaikan variablenya
	*/

	for _, val := range hari {
		fmt.Println("hari", val)
	}
	fmt.Println()

	// untuk for pada slice sama saja ya seperti array

	// for range map

	var orang map[string]string = map[string]string{
		"nama":    "muiz",
		"jurusan": "teknik informatika",
	}

	for key, val := range orang {
		fmt.Println(key, "=", val)
	}

	/*

		akan tetapi untuk map variable pertama setelah for itu bukan index
		melaikan keynya

		jadi seperti itu for loop digolang
		mudah mudahan kalian paham

	*/

}
