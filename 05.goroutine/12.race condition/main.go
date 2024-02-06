/*

- masalah dengan goroutine

saat kit amenggunakan gorotuine, dia tidak hanya berjalan secara concurrency, tetapi bisa secara
paraller juga

karena bisa ada beberapa thread yang berjalan secara paraller

hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh
beberapa goroutine secara bersamaan

jadi ada satu buah variable yang diakses oleh beberapa goroutine
ini sangat berbahaya kerena akan menyebabkan race condition

didalam bahasa pemrograman yang memiliki fitur paraller akan mengalami masalah tersebut
dan tenang saja dimateri berikutnya kita akan coba menangani hal tersebut

untuk sekarang kita coba pahami dulu tentang race condition
oke langsung saja kita simulasikan race conditionnya




*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var x = 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(x)

	/*

		jadi disini saya membuat gorotine sebanyak 1000 kali yang masing masing gorotine tersebut
		menjalankan looping untuk incrementnya sebanyak 100

		jadi total angka yang harusnya kita dapatkan dari variable x itu adalah 100000
		akan tetapi kalo kita jalankan hasilnya tidak tentu

		91308
		89466

		kenapa bisa seperti ini ?

		golang itu threadnya berjalan secara paraller, jadi kita tidak tahu gorotinenya ada
		dithread mana saja


		ada kalanya ketika sebuah variable x itu bernilai 100, maka x akan disharing ke beberapa goroutine
		dan kebetulan gorotine tersebut berjalan secara paraller pada thread yang berbeda

		gorotine1 x = x + 1
		gorotine2 x = x + 1
		gorotine2 x = x + 1

		karena dia berjalan secara bersamaan maka nilai xnya adalah 101 bukan 103
		karena dari ketiga gorotine tersebut nilai x adalah 100

		jadi inilah yang dinamakan dengan race condition
		dimateri selanjutnya kita akan coba mengatasi hal tersebut





	*/

}
