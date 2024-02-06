/*

- break dan continue

break dan continue adalah keyword yang dapat kita gunakan didalam perulangan
break bisa kita gunakan untuk menghentikan paksa perulangan dan
continue bisa kita gunakan untuk melakukan skip perulangan dan melakukan ke
perulangan selanjutnya

oke langsung saja kita coba

*/

package main

import "fmt"

func main() {

	// break

	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}

		fmt.Println("perulangan ke", i)
	}
	fmt.Println("perulangan selesai")
	fmt.Println()

	/*

		kalo kita jalankan maka dia akan mulai dari angka 0 sampai ke 6
		karena ketika variable i sudah mencapai angka 6 akan dicheck

		apakah i lebih kecil 10 ? true
		masuk kedalam perulangan

		jika i  sama dengan 6 maka jalankan break atau hentikan perulangannya
		maka code yang dibawahnya tidak akan dijalankan

	*/

	// continue

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ke", i)
	}
	fmt.Println("perulangan selesai")

	/*

		kalo kita jalankan maka hasilnya akan seperti ini

		perulangan ke 1
		perulangan ke 3
		perulangan ke 5
		perulangan ke 7
		perulangan ke 9
		perulangan selesai

		jadi kalo variable i itu bernilai genap maka jalankan continue
		dan code yang dibawahnya tidak akan dijalankan

		tetapi pengulangannya tetap akan berjalan
		jadi seperti itu ya kata kunci break dan continue

		mudah mudahan kalian paham

	*/

}
