/*


- package os


nah sekarang kita akan mulai mencoba package package apa saja yang telah
disediakan oleh golang

kita coba package os terlebih dahulu
jadi package os ini kita bisa gunakan untuk mengakses fitur sistem operasi secara independen
jadi kita tidak perlu memikirkan sistem operasinya apa, program pada package tersebut
akan sama saja

oke langsung saja kita coba

*/

package main

import (
	"fmt"
	"os"
)

func main() {

	var result []string = os.Args
	fmt.Println(result[1:])

	/*

		jadi os.Args itu digunakan untuk mengambil argument,
		jadi ketika kalian menjalankan programnya menggunakan perintah

		go run main.go muiz

		setelah nama filenya kalian bisa menambahkan argument
		dalam kasus ini adalah muiz

		dan result datanya adalah slice artinya kalian bisa menambahkan argument
		sebanyak apapun ya

		nah selanjutnya kita bisa mengambil hostName
		jadi hostname itu adalah nama yang digunakan untuk membedakan antara satu perangkat
		dengan perangkat yang lain pada jaringan internet

		kita coba ya

	*/

	var hostName, err = os.Hostname()
	if err == nil {
		fmt.Println(hostName)
	} else {
		fmt.Println(err.Error())
	}

	/*

		jadi Hostname mengembalikan 2 return value, ada hostnamenya dan error
		sebenarnya ada banyak sekali ya, kalian bisa baca baca disini

		https://pkg.go.dev/os

	*/

}
