/*

- interface kosong

kali ini kita akan belajar mengenai interface kosong

apa itu interface kosong ?
interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara
otomatis semua type data akan menjadi implementasinya

atau simplenya nih, interface adalah sebuah type data yang bisa menampung segala jenis data
jadi kalo kalian punya variable yang memiliki type data interface kosong

maka variable tersebut bisa kita assigment dengan type data apapun
dan interface kosong ini memiliki alias any

jadi sebenarnya secara tidak langsung kita sudah menggunakan interface kosong
secara tidak langsung

ketika kalian memunculkan value ke terminal kita bisa gunakan fmt.Println()
dan kita bisa memasukkan apapun didalam parameter Println tersebut

artinya function tersebut menggunakan type data interface kosong sebagai
parameternya

ada juga panic, recover juga menggunakan interface kosong sebagai parameternya
dan lain lain

oke langsung saja kita coba


*/

package main

import "fmt"

func dynamicType(value any) any {
	return value
}

func main() {

	var value interface{} = 90
	value = "apapun"
	fmt.Println(value)

	/*

		jadi kalo misalkan kita buat variable dengan type data interface kosong
		maka kita bisa memasukka type data apapun kedalam variable tersebut

		atau juga kedalam parameter dan return value dari function
		kita coba

	*/

	fmt.Println(dynamicType("halo"))
	fmt.Println(dynamicType(true))
	fmt.Println(dynamicType(10))

	/*

		jadi seperti itu ya interface kosong, jadi kita punya kontrak dengan interface kosong
		ketika didalam kontrak tersebut tidak ada definisi method, maka semua type data
		bisa mengimplementasikan interface kosong

		akan tetapi dari type data any atau inteface kosong
		kita tidak bisa secara langsung memasukkan kedalam sebuah variable
		yang memiliki type data selain any

		misalkan seperti ini

	*/

	// var nama string = dynamicType("muiz")
	// fmt.Println(nama)

	/*

		kalo kita jalankan maka akan terjadi error, nah kenapa seperti ini?
		padahan kan yang kita masukkan ada string dan function tersebut
		juga mengembalikan string

		kenapa error? karena meskipun data yang kita lihat itu bertype string
		akan tetapi sebenarnyakan yang dikembalikan bertype any

		kalo kita jaankan maka errornya seperti ini
		cannot use dynamicType("muiz") (value of type any) as string value in variable declaration: need type assertion

		katanya tidak bisa memasukkan value dari type data any sebagai type data string
		pada deklarasi variable, dibutuhkan type assertion

		nah kalo kalian ingn mengconvert dari type data any ke type data
		yang diinginkan maka kalian butuh type assertion

		apa itu type assertion?
		nanti akan kita bahas dimateri berikutnya

		oke jadi seperti itu ya tentang interface kosong atau any
		mudah mudahan kalian paham



	*/
}
