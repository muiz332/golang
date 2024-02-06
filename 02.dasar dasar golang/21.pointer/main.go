/*

- pointer


dimater kali ini kita akan membahas mengenai pointer

secara default digolang itu pas by value bukan by reference
artinya, jika kita mengirimkan sebuah variable kedalam sebuah function, method atau variable lain
sebenarnya yang dikirim adalah duplikasi valuenya, karena by default adalah pas by value

nah sekarang kita coba

*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{
		City:     "banyuwangi",
		Province: "jawa timur",
		Country:  "indonesia",
	}

	address2 := address1

	address2.City = "malang" // ubah properti city

	fmt.Println(address1)
	fmt.Println(address2)

	/*

		kalo kita jalankan maka hasilnya seperti ini

		{banyuwangi jawa timur indonesia} // address1
		{malang jawa timur indonesia} // address2

		address1 citynya tetap banyuwangi sedangkan yang berubah hanya addres2
		citynya menjadi malang

		mungkin ini terlihat biasa saja, akan tetapi kalo kalian berasal dari bahasa pemrograman
		yang lain misalkan javascript

		didalam javascript kalo kita membuat sebuah object kemudian kita assigment ke address1
		dan kita buat address2 yang kita assigment dengan address1

		kemudian address2 kita ubah value dari propertynya
		maka address1 akan ikut berubah, jadi didalam memory address1 dan address2 itu
		lokasinya sama yang ada didalam address2 adalah alamat memorynya saja

		inilah yang disebut dengan pas by reference

		akan tetapi digolang tidak begitu
		digolang itu pas by value, artinya ketika variable address1 kita assigment ke variable
		address2 maka yang terjadi adalah

		address1 akan dicopy dan diassigment kedalam variable address2
		jadi didalam memory address1 dan address2 itu memiliki lokasi memory yang berbeda

		jadi itulah pas by value



		pointer itu adalah kemampuan membuat reference kelokasi data dimemory yang sama
		tanpa menduplikasi data yang sama

		jadi simplenya pointer itu adalah kemampuan untuk membuat pas by reference

		lalu bagaimana kalo kita ingin membuat pointer digolang?
		kita bisa menggunakan operator & diikuti dengan nama variablenya

		kita coba

	*/

	var address3 Address = Address{
		City:     "banyuwangi",
		Province: "jawa timur",
		Country:  "indonesia",
	}

	var address4 *Address = &address3

	address4.City = "malang" // ubah properti city

	fmt.Println(address3)
	fmt.Println(address4)

	/*

		jadi kalo kalian ingin variable address4 itu datanya mengacu ke address3
		maka kalian bisa menambahkan operator & diikuti dengan nama variablenya

		kalo kalian jalankan maka hasilnya akan seperti ini
		{malang jawa timur indonesia}
		&{malang jawa timur indonesia}

		pada property city kedua duanya itu berubah datanya dengan nilai yang sama
		yaitu malang

		dan kalo kalian perhatikan, data pointer itu ada operator & didepan
		datanya

		ini menunjukkan bahwa data tersebut adalah sebuah pointer

	*/
}
