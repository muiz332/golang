/*

- hati hati saat membuat array

hati hati saat membuat array, bisa jadi kalo kita salah menulis
yang kita buat bukan array tetapi slice

contoh

*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// ini array
	var iniArray1 [10]string = [10]string{
		"coba",
		"coba",
	}

	// ini array
	var iniArray2 = [...]string{"senin", "selasa"}

	// ini slice

	var iniSlice = []string{"rabu", "kamis"}

	fmt.Println(reflect.ValueOf(iniArray1).Kind())
	fmt.Println(reflect.ValueOf(iniArray2).Kind())
	fmt.Println(reflect.ValueOf(iniSlice).Kind())

	/*

		kalo kita lihat type datanya hasilnya akan seperti ini

		array
		array
		slice

		jadi kalo kalian tidak menuliskan panjang arraynya
		itu akan dianggap sebagai slice

		jika kalian tidak tahu berapa panjang array yang mau dibuat
		maka tambahkan titik 3 kali didalam kurung kotaknya

		jadi untuk membuat slice itu ada 2 cara
		cara pertama menggunakan function make

		dan cara yang kedua seperti array, tetapi kita tidak menuliskan
		ponjang atau titik 3 kali didalam kurung kotak

		jadi hanya kurung kotak saja
		dan digolang nanti kebanyakan yang akan kita pakai
		adalah type data slice ya dibanding type data array

		salah satu alasannya karena slice itu bisa agak sedikit dynamic
		walaupun ada capacitynya

		jadi sampai disini dulu materi kita tentang slice
		mudah mudahan kalian paham

	*/

}
