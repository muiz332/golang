/*

- pointer difuncttion

secara default digolang yang menerapkan kemampuan pas by reference salah satunya
adalah slice, map

struct secara default menggunakan pas by value
jadi ketika kalian mengirimkan struct sebagai parameter, kalian sebenarnya
mengcopy data structnya

kita bisa menggunakan pointer agar structnya menjadi pas by reference

kita coba yang pass by value dulu ya

*/

package main

import "fmt"

type Mahasiswa struct {
	nama     string
	jurusan  string
	angkatan string
}

func sayHelloValue(mhs Mahasiswa) {
	mhs.nama = "hallo"
}

func sayHelloReference(mhs *Mahasiswa) {
	mhs.nama = "hallo"
}

func main() {

	var hasan Mahasiswa = Mahasiswa{
		nama:     "hasan",
		jurusan:  "agama",
		angkatan: "2021",
	}

	sayHelloValue(hasan)
	fmt.Println(hasan) // {hasan agama 2021}

	/*

		kalo saya jalankan, maka tidak akan berubah property dari namanya ya
		nah sekarang kita coba terapkan pointer pada parameternya

	*/

	var husain Mahasiswa = Mahasiswa{
		nama:     "husain",
		jurusan:  "filsafat",
		angkatan: "2020",
	}

	sayHelloReference(&husain)
	fmt.Println(husain) // {hallo filsafat 2020}

	/*

		kalo saya jalankan maka properti nama akan berubah ya
		jadi itulah pointer difunction

		nanti kalian akan sering memnggunakan pointer difunction ini
		khusunya ketika berinteraksi dengan struct

		mudah mudahan kalian paham

	*/
}
