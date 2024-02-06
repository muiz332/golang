/*

- type data array

array adalah sebuah kumpulan data dengan type data yang sama dimana banyak jumlah
kumpulan tersebut harus ditentukan dulu berapa jumlahnya
jadi simplenya array itu lebih sakti dari pada variable

kalo variable hanya bisa menyimpan satu nilai, maka array bisa menyimpan
banyak nilai

langsung saja kita coba

*/

package main

import "fmt"

func main() {
	// cara 1

	var bulan [3]string
	bulan[0] = "january"
	bulan[1] = "february"
	bulan[2] = "maret"

	fmt.Println(bulan)

	/*

		jadi kita buat variablenya dengan nama bulan, dimana bulan memiliki type
		data seperti ini [3]string

		artinya saya berikan type data array dimana array tersebut memiliki data yang bertype
		string, dengan jumlahnya maximal 3 data

		untuk mengakses array kita gunakan index
		jadi kalian bisa tulis seperti ini

		bulan[0]
		maka yang muncul adalah january

		kalo kalian assigment ulang misalkah

		bulan[0] = "desember"
		maka bulan dengan index atau urutan ke 0 itu valuenya desember
		jadi itu digunakan untuk mengubah data didalam array

	*/

	// cara 2

	var hari [3]string = [3]string{"senin", "selasa"}
	fmt.Println(hari[0])
	fmt.Println(hari[1])

	/*

		jadi kita langsung assigment dengan data arraynya
		kalian gunakan kurung kurawal buka dan tutup
		kemudian didalamnya kalian tuliskan stringnya

		dan pisahkan dengan koma jika lebih dari satu

		didalam array, kalo kalian mengakses datanya melebihi
		jumlah datanya, maka akan error

		index out of bounds
	*/

	// fmt.Println(hari[3]) // error

	// function pada array

	/*

		kita bisa gunakan len untuk memerikan panjang arraynya
		ingat panjang arraynya bukan berapa banyak datanya
	*/

	fmt.Println(len(hari))

}
