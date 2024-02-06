/*

- context with deadline


nah kita juga bisa menambahkan context with deadline, sebenarnya sama saja seperti
context with timeout akan tetapi kita bisa menentukan kapan sinyal cancel itu dikirim
secara spesifik

contohnya nanti jam 1 saya ingin mengirimkan sinyal cancel, nah itu kita gunakan deadline
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
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) // cancel with timeout
	defer cancel()

	var result = goroutineLeakFix(ctx)

	for data := range result {
		fmt.Println(data)
	}

	fmt.Println("jumlah goroutine akhir", runtime.NumGoroutine())

	/*

		jadi seperti itu ya context with deadline, jadi kita harus menentukan kapan
		waktu timeoutnya berakhir

		jadi kita bisa gunakan time.Now atau waktu sekarang ditambah dengan 5 detik
		jadi seperti itu ya penggunaan dari deadline

		mudah mudahan kalian paham

	*/

}
