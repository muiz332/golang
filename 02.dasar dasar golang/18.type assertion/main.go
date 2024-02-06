/*

- type assertion


type assertion adalah kemampuan untuk mengubah type data any menjadi type data
yang diinginkan

jadi type assertion sering digunakan ketika bertemu dengan type data any

oke langsung saja kita coa


*/

package main

import (
	"fmt"
	"reflect"
)

func apapun(num int) any {
	if num == 0 {
		return "hello"
	} else if num == 1 {
		return true
	}
	return 10
}

func main() {

	var result = apapun(0)
	var ubahString string = result.(string)
	fmt.Println(ubahString)

	/*

		jadi tinggal kita tambahkan titik setelah variablenya kemudian
		diikuti dengan kurung buka dan tutup yang didalamnya terdapat
		type data apa yang kita inginkan

		kalo kita jalankan  maka tidak error

		akan tetapi kalo kalian salah menuliskan type datanya
		misalkan returnnya itu boolean kalian convert ke string

		maka akan terjadi panic dan panicnya tidak ter recover, maka program akan berhenti
		jadi kalian harus berhati hati ya jangan sampai salah convert

		sebenarnya kita bisa gunakan switch untuk menanggulagi panic
		kita coba

	*/

	var hasil = apapun(1)

	switch hasil.(type) {
	case string:
		fmt.Println(hasil, "bertype = ", reflect.ValueOf(hasil).Kind())
	case bool:
		fmt.Println(hasil, "bertype =", reflect.ValueOf(hasil).Kind())
	default:
		fmt.Println(hasil, "bertype =", reflect.ValueOf(hasil).Kind())
	}

	/*
		jadi kita bisa check nih menggunakan switch, dengan menambahkan kondisi
		variablenya titik kemudian kurung buka dan tutup dan didalamnya
		terdapat type, dimana type ini hanya bisa digunakan pada switch ya

		kita bisa check menggunakan case
		kalo type datanya string kita mau apain

		kalo type datanya boolean mau diapain, atau selain itu kita mau apain
		jadi menggunakan switch akan lebih aman dan menghindari panic

		jadi seperti itu type assertion
		mudah mudahan kalian paham


	*/
}
