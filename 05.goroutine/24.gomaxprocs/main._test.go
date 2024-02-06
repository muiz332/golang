/*

- gomaxprosc


sebelumnya kita sudah bahas bahwa goroutine itu berjalan didalam thread
pertanyaannya berapa banyak thread yang ada digolang?

untuk mengetahui jumlah thread kita bisa gunakan gomaxprocs
yaitu sebuah function didalam package runtime yang bisa kita gunakan untuk
mengubah jumlah thread atau melihat jumlah thread

secara default golang akan membuat thread sebanyak cpu dikomputer kita
untuk melihat jumlah cpu kita bisa gunakan runtime.Numcpu()

langsung saja kita coba

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestTotalCpu(t *testing.T) {
	var total int = runtime.NumCPU()
	fmt.Println("total cpu", total)
}

func TestTotalThread(t *testing.T) {
	var total = runtime.GOMAXPROCS(-1) // kalo diatas -1 itu mengubah jumlah threadnya
	fmt.Println("total thread", total)
}

func TestTotalGoroutine(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println("goroutine", num)
		}(i)
	}

	var total int = runtime.NumGoroutine()
	wg.Wait()
	fmt.Println("total goroutine berjalan", total)

	/*

		kenapa mucul 12? bukannya 10?
		karena secara default golang akan menjalankan 2 buah goroutine dibelakang layar

		nah yang 10 itu adalah goroutine yang sudah kita buat sendiri ya

		jadi itu lah gomaxprocs, jadi sebenarnya kalian tidak perlu mengubah jumlah
		threadnya semuanya sudah diatur oleh go scaduler

		oke jadi seperti itu ya
		mudah muahan kalian paham

	*/
}
