/*

-  pointer dimethod

walaupun method itu menempel pada structnya, akan tetapi secara default
sebenarnya struct yang kita akses didalam function itu adalah duplikat
dari structnya, artinya pas by value juga

sangat direkomendasikan ketika kalian memiliki method di structnya
kalian menerapkan pointer

sekarang kita coba ya

*/

package main

import "fmt"

type Orang struct {
	nama string
	umur uint
}

func (orang Orang) sayHelloValue() {
	orang.nama = "hallo nama saya" + orang.nama
}

func (orang *Orang) sayHelloReference() {
	orang.nama = "hallo nama saya " + orang.nama
}

func main() {
	var hasan Orang = Orang{
		nama: "hasan",
		umur: 10,
	}

	hasan.sayHelloValue()
	fmt.Println(hasan) // {hasan 10}

	/*

		kalo kita tidak menggunakan pointer maka
		ketika kita ubah datanya maka dia tidak akan berubah

		ketika nanti kalian memiliki struct yang banyak
		dan struct itu diuplikasi terus ketika memanggil methodnya
		maka itu akan sangat memakan memory

		hal itulah kenapa kita wajib mengimplementasikan pointer dimethod
		sekarang kita coba method yang sudah menerapkan pointer

	*/

	hasan.sayHelloReference()
	fmt.Println(hasan) // {hallo nama saya hasan 10}

	/*

		jadi kalian tinggal tambahkan tandah * untuk memberitahu sigolang bahwa
		parameter ini akan menerapkan pointer

		jadi seperti itu ya pointer dimethod
		mudah mudahan kalian paham

	*/

}
