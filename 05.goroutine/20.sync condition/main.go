/*


- sync cond

cond adalah implementasi locking berbasis kondisi

cond membutuhkan locker jadi kita bisa menggunakan mutex atau rwmutex untuk implementasi
lockingnya

namun berbeda dengan locker biasanya, dicond terdapat function wait untuk menunggu apakah
perlu menunggu atau tidak

function disignal bisa kita gunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu
lagi, tetapi kita harus panggil signal sebanyak goroutine yang telah kita buat

sedangkan function broadcast digunakan untuk memberitahu semua goroutine agar tidak
perlu menunggu lagi

untuk membuat cond kita bisa menggunakan function sync.NewCond(locker)

oke kita coba

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitCondition(value int, cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	cond.L.Lock()
	cond.Wait() // ditunggu sampai ada yang mengirim signal baru dia jalan ke goroutine selanjutnya
	fmt.Println("done", value)
	cond.L.Unlock()

}

func main() {

	var cond = sync.NewCond(&sync.Mutex{})
	var wg = &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go WaitCondition(i, cond, wg)
	}

	go func() {
		for i := 0; i < 10; i++ {

			time.Sleep(1 * time.Second)

			cond.Signal()
		}

	}()
	wg.Wait()

	/*

		ketika pertama kali dijalankan maka goroutinenya akan melakukan locking dulu
		setelah locking akan melakukan wait yaitu menunggu sampai ada signal
		kemudian diunlock

		lalu bagaimana cara mengirim signalnya ?
		kita bisa gunakan looping dan misalkan kita ingin mengirim signalnya satu detik sekali
		jadi seperti itu ya

		maka nanti munculkan akan secara bergantian

		nah selanjutnya kita bisa pakai broadcast agar ketika menunggu akan langsung
		dijalankan

		jadi kita bisa tulis sepert ini

	*/

	// go func() {
	// 	cond.Broadcast()
	// }()

	// wg.Wait()

	/*

		jadi untuk mengirim signal atau broadcastnya itu kalian harus gunakan goroutine ya
		kalo tidak akan akan terjadi error deadlock

		jadi seperti itu ya sync.cond

		jadi kita bisa mengatur kapan goroutine itu harus menunggu dan kapan
		goroutine itu harus berjalan, tentu saja berjalan dengan aman tanpa race condition ya

		kalo kalian butuh locking yang ingin diatur kapan jalannya, kalian bisa gunakan
		sync.newCond ya

		oke jadi seperti itu mudah muahan kalian paham

	*/

}
