/*


- struct method


struct itu sebenarnya sama dengan type data yang lainnya, dia bisa digunakan untuk parameter
pada function

akan tetapi kalian bisa menambahkan sebuah function pada structnya seolah oleh struct tersebut
memiliki function

seperti di bahasa pemrograman OOP ya, class itu memiliki method, nah kalo digolang juga bisa seperti itu
walupun dibelakang layar sebenarnya bukan method, tetapi function terpisah

karena kita tidak bisa mendeklarasikan sebuah method didalam struct

langsung saja kita coba

*/

package main

import "fmt"

type Orang struct {
	Nama    string
	Address string
	Umur    uint8
}

func (orang Orang) sayHello() {
	fmt.Println("halo nama saya", orang.Nama)
}

func salam(nama string, orang Orang) {
	fmt.Println("hello", nama, "perkenalkan nama saya", orang.Nama)
}

func main() {
	var muiz Orang = Orang{
		Nama:    "muiz",
		Address: "banyuwangi",
		Umur:    10,
	}

	muiz.sayHello()

	/*

		jadi untuk membuat struct method kalian bisa buat terlebih dahulu structnya
		dan kita bisa buat methodnya diluar deklarasi structnya

		misalkan saya buat namanya sayHello, akan tetapi sebelum nama methodnya
		saya harus menuliskan kurung buka dan tutup

		yang didalamnya terdapat type data dari structnya untuk memberitahukan
		bahwa, function sayHello ini milik struct apa

		kita bisa tulis tuh namanya orang dari struct Orang, untuk namanyanya bebas saja
		sebenarnya tidak apa apa, tetapi yang paling penting adalah dari struct yang mana

		ketika sudah seperti ini, maka kalian bisa panggil dengan cara menuliskan nama
		struct yang sudah diberikan value, atau istilahnya initialisasi, kedalam variable muiz

		karena digolang tidak akan konsep this, maka untuk mengakses structnya
		seperti ini (orang Orang), jadi tinggal kalian tulis orang.Nama
		kalo ingin mengakses nama

		jadi untuk memanggilnya kalian bisa tulis

		muiz.sayHello()

		kalo kalian hanya menuliskan nama functionnya saja, maka itu akan terjadi error

		nah meskipun kelihatannya struct itu memiliki method, akan tetapi sebenarnya yang
		terjadi dibelakang layar methodnya terpisah

		jadi bukan method melainkan function biasa
		jadi kalo kita buat yang sebenarnya terjadi seperti ini
	*/

	salam("hasan", muiz) // hello hasan perkenalkan nama saya muiz

	/*

		jadi sebenarnya structnya kita masukkan kedalam sebuah parameter seperti ini
		lebih tepatnya function ya bukan method

		jadi kalo dibelakang layar yang terjadi seperti function salam
		bedanya, ketika membuatnya kurungnya sebelum nama functionnya
		dan ketika memanggilnya harus menggunakan nama struct yang
		telah kita assigment atau initilisasi

	*/

}
