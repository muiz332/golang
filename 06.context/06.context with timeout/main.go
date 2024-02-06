/*

- context with timeout

selain menambahkan value kedala context dan juga sinyal cancel, kita juga menambahkan sinyal
cancel secara otomatis dengan menggunakan timeout

dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara
manual, cancel akan secara otomatis dieksekusi jika waktu timeout sudah terlewati

penggunaan timeout sangat cocok ketika kita melakukan query kedatabase atau htttp api
namun ingin menentukan batal maksimal timeoutnya

untuk membuat context dengan cancel sinyal secara otomatis menggunakan timeout
kita bisa gunakan function

context.WithTimeout(parent, duration)

oke langsung saja kita coba





*/

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func goroutineLeakFix(ctx context.Context) chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				channel <- counter
				counter++
				time.Sleep(1 * time.Second) // simulasi slow proses

			}
		}

	}()

	return channel
}

func main() {

	fmt.Println("jumlah goroutine awal", runtime.NumGoroutine())

	var parent = context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second) // cancel with timeout
	defer cancel()

	var result = goroutineLeakFix(ctx)

	for data := range result {
		fmt.Println(data)
	}

	fmt.Println("jumlah goroutine akhir", runtime.NumGoroutine())

	/*

		nah kaian tinggal buat seperti itu ya, jadi mirip seperti cancel ya
		bedanya kita bisa menambhakan parameter timeoutnya

		dan kita tetap harus memanggul function cancelnya pada defer
		kenapa seperti itu?

		jadi kan timeout itu adalah kejadian pembatalan proses ketika waktunya melebihi batas
		yang kita berikan

		nah kalo misalkan waktunya belum melebihi dari waktu yang kita berikan
		maka kita tetap harus memanggil cancel

		karena biar tidak ada proses dibackground yang masih berjalan
		karena kan prosesnya cepat, jadi karena prosesnya cepat

		kita tidak perlu menunggu melebihi waktu timeoutnya
		jadi kita bisa langsung cancel saja untuk menghindari goroutineleak

		jadi seperti itu ya
		jadi kita tetap gunakan defer untuk tetap mengeksekusi cancelnya
		diakhir function tersebut

		dan perlu kalian ingat kalo kalian menggunakan timeout maka tidak perlu
		menggunakan wait grub ya

		oke jadi seperti itu ya
		mudah mudahan kalian paham

	*/
}
