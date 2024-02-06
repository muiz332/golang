/*

- atomic

sebelumnya ketika kita punya sebuah variable yang kita sharing kebanyak goroutine
dan kita juga ingin mengatasi race condition

kita bisa menggunakan sync.Mutex

nah sebenarnya ketika kita ingin mengatasi masalah race condition pada type data primitif
contohnya int, string, boolean, float

maka kita bisa gunakan atomic

oke langsung saja ya kita coba


*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var angka int64 = 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&angka, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("hasilnya", angka)
	/*

		jadi wg.add itu kalian panggil sebelum goroutinenya dijalankan ya
		Hal ini untuk memastikan bahwa WaitGroup (wg) mengetahui bahwa akan ada goroutine tambahan yang akan dijalankan dan bahwa program tidak menunggu untuk menyelesaikan goroutine yang tidak akan pernah dimulai.

		jadi itulah atomic, jadi kalian tidak perlu menggunakan mutex ya
		tetapi kalo berhubungan dengan struct kalian tetap harus menggunakan mutex

		oke jadi seperti itu ya
		mudah mudahan kalian paham

	*/

}
