/*

- pointer operator *


jadi tanda * itu bisa digunakan sebagai type data dari sebuah pointer
dan juga bisa digunakan untuk operator

jadi kalo kita gunakan tanda * sebagai operator kita bisa mengubah seluruh


*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var address1 Address = Address{
		City:     "Banyuwangi",
		Province: "Jawa timur",
		Country:  "Indonesia",
	}

	var address2 Address = Address{
		City:     "Jogja",
		Province: "Jawa tengah",
		Country:  "Indonesia",
	}

	var address3 *Address = &address1

	address3 = &address2

	fmt.Println(address1) // {Banyuwangi Jawa timur Indonesia}
	fmt.Println(address2) // {Jogja Jawa tengah Indonesia}
	fmt.Println(address3) // &{Jogja Jawa tengah Indonesia}
	fmt.Println()

	/*

		kalo kalian jalankan ketika saya mengubah value dari variablenya maka
		pointernya akan berpindah

		yang awalnya dari address1 akan berubah ke address2
		kalo kita ubah address3nya seperti ini

	*/

	address3.City = "Madiun"
	fmt.Println(address1) // {Banyuwangi Jawa timur Indonesia}
	fmt.Println(address2) // {Madiun Jawa tengah Indonesia}
	fmt.Println(address3) // &{Madiun Jawa tengah Indonesia}
	fmt.Println()

	/*

		lalu bagaimana kalo saya ingin ketika address3 saya assigment dengan
		address1 maka address2 akan ikut berubah

		kita bisa gunakan operator *
		kita coba

	*/

	*address3 = address1

	fmt.Println(address1) // {Banyuwangi Jawa timur Indonesia}
	fmt.Println(address2) // {Banyuwangi Jawa timur Indonesia}
	fmt.Println(address3) // &{Banyuwangi Jawa timur Indonesia}
	fmt.Println()

	/*

		kalo kita jalankan maka value dari address2 akan menjadi address1
		tetapi pointernya akan tetap ke address2, bukan ke addres1

		jadi variable apapun yang memiliki reference ke address2
		maka nilainya juga akan berubah

		kita coba mengubah citynya

	*/

	address3.City = "Malang"

	fmt.Println(address1) // {Banyuwangi Jawa timur Indonesia}
	fmt.Println(address2) // {Malang Jawa timur Indonesia}
	fmt.Println(address3) // &{Malang Jawa timur Indonesia}
	fmt.Println()

	/*

		kalo kita jalankan yang berubah hanya address2 dan addreess3
		sedangkan address1 nilainya tetap ya

		jadi pointernya tetap ke address2


	*/

}
