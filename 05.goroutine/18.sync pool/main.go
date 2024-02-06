/*


- sync.pool


pool adalah implementasi design pattern bernama object pool pattern
sederhananya pool ini digunakan untuk menyimpan data

jadi anggap saja kita buat tempat penyimpanan data yaitu pool atau kolam
lalu untuk menggunakan datanya kita bisa mengambil data dari pool

setelah selesai data digunakan maka datanya harus dikembalikan dipoolnya
bisanya dalam kenyataan implementasi pool ini sering sekali digunakan untuk menanage
data connection ke database

jadi daripada kita setiap ingin terkoneksi dengan database kita membuat koneksi baru
lebih baik diawal ketika applikasinya menyala kita akan bikin misalnya 10
koneksi yang disimpan dipool

nah kalo kita butuh koneksi kita hanya perlu mengambilnya dari pool tadi
dan setelah selesai kita akan kembalikan lagi kepoolnya

jadi ini adalah object pool pattern
implementasi pool digolang itu sudah aman dari race condition
jadi kalian tidak perlu menggunakan mutex ya untuk menangani hal tersebut

oke langsung saja kita coba


*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}
	var pool sync.Pool

	// var pool = sync.Pool{
	// 	New: func() any {
	// 		return "nilai default"
	// 	},
	// }

	pool.Put("hasan")
	pool.Put("husain")
	pool.Put("muiz")

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			data := pool.Get() // mengambil data secara random pada pool
			fmt.Println(data)
			pool.Put(data) // kita harus mengembalikan data ke pool lagi
		}()
	}

	wg.Wait()
	fmt.Println("selesai")

	/*

		kalo kita jalankan maka hasilnya seperti ini

		husain
		hasan
		<nil>
		<nil>
		<nil>
		<nil>
		<nil>
		<nil>
		<nil>
		muiz
		selesai

		kenapa ada data yang nil ? karena belum sempat datanya dikembalikan
		goroutinenya terlalu cepat untuk mengambil datanya

		lalu kalo misalkan kalian ingin menggunakan default value
		yang tadinya nil kalian bisa gunakan value yang lain

		itu kalian gunakan key New seperti ini

		seperti ini

		var pool = sync.Pool{
			New: func() any {
				return "nilai default"
			},
		}

		jadi nanti kalian bisa merubah nilai nil jadi value nilai default

		dan satu lagi


		kalian menuliskan seperti ini
		var pool sync.pool

		itu sama saja dengan kalian menuliskan seperti ini

		var pool = sync.pool{}

		bedanya cara pertama itu kalian tidak ingin memasukkan properti didalam structnya
		dan cara kedua itu kalo kalian ingin memasukkan properti didalam structnya

		makanya kalo ingin menggunakan proprti new kalian harus menggunakan
		cara yang ke 2

		jadi seperti itu ya mudah mudahan kalian paham


	*/

}
