/*

- nil

sekarang kita akan belajar mengenai nil

biasanya pada bahasa pemrograman yang lain, sebuah variable atau object yang belum kita
inisialisasi maka secara otomatis nilainya adalah null atau nil

berbeda dengan golang, saat kita membuat variable dengan type data tertentu
maka secara otomatis akan dibuatkan default valuenya mengikuti type datanya ap

nil itu bisa kita gunakan sebagai data kosong
jadi kalo misalkan sebuah variable kita isi dengan nil maka dianggap sebagai
variable yang memiliki data kosong

akan tetapi nil hanya bisa digunakan dibeberapa type data saja seperti
interface, function,map,slice, pointer dan channel

untuk pointer dan channel kita akan bahas digoroutine





*/

package main

import "fmt"

func iniMap(value string) map[string]string {
	if value == "" {
		return nil
	}

	return map[string]string{
		"satu": "january",
	}
}

func main() {
	var bulan map[string]string = nil
	fmt.Println(bulan)

	// sebuah function yang mereturn
	var result map[string]string = iniMap("")
	fmt.Println(result)

	if result == nil {
		fmt.Println("ini map kosong")
	}

	/*

		jadi seperti itu ya penggunaan nil digolang
		nil ini akan sering kalian gunakan saat bertemu dengan interface error

		jadi seperti itu mudah mudahan kalian paham

	*/
}
