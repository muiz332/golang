/*

- sync wait grub


sekarang kita akan bahas mengenai waitgrub, jadi wait grub ini adalah fitur untuk menunggu sebuah proses
selesai dilakukan

hal ini kadang diperlukan ketika misalkan kita ingin menjalankan beberapa proses goroutine
tetapi ingin menunggu semua proses goroutine tersebut selesai

kasus seperti ini bisa menggunakan wait grub, sebelumnya kita hanya menggunakan time.Sleep untuk
menunggu goroutinenya selesai tetapi kita tidak tahu kapan selesainya

untuk menandai bahwa ada proses goroutine kita menggunakan method Add(int), setelah proses
goroutinenya selesai kita menggunakan method done

untuk menunggu semua proses selesai kita menggunakan wait

oke langsung saja kita coba ya





*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func RunAsync(group *sync.WaitGroup, angka int) {
	defer group.Done()
	group.Add(1) // 1 untuk 1 goroutine

	time.Sleep(2 * time.Second)
	fmt.Println("hello", angka)
}

func main() {

	group := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go RunAsync(group, i)
	}

	group.Wait()
	fmt.Println("semua goroutine selesai")

	/*

		jadi ini mirip dengan channel (alat komunikasi secara syncronous) tetapi kalo channel
		kita dapat menunggu sebuah goroutine selesai dijalankan dan mendapatkan nilainya

		kalo kita jalankan hasilnya seperti ini

		hello 0
		hello 2
		hello 4
		hello 3
		hello 1
		semua goroutine selesai

		kalo waitgrub hanya untuk menunggu semua proses goroutine yang kita punya selesai dijalankan
		jadi itulah waitGroup

		method add itu sebenarnya hanya digunakan untuk mengetahui jumlah goroutine
		nah untuk method done itu digunakan untuk mengurani jumlah goroutine tadi

		karena sudah selesai dijalankan
		dan wait untuk menunggu prosess done tersebut menjadi 0

		mudah mudahan kalian paham ya

	*/

}
