/*

-  membuat channel


channel digolang direpresentasikan dengan type data chan
untuk membuat channel itu sangat mudah kita cukup menggunakan function make()
mirip saat kita membuat map atau slice

tapi dalam pembuatan channel, kita harus menentukan type data apa yang bisa kita masukkan
kedalam channel tersebut

jadi kalian bisa tulis sepertini

channel := make(chan string)

langsung saja kita coba ya




*/

package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	var channel = make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "ini dari database"
		fmt.Println("selasai menjalankan function goroutine")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

	/*

		jadi kalo kita jalankan maka hasilnya akan menjadi seperti ini

		=== RUN   TestCreateChannel
		selasai menjalankan function goroutine
		ini dari database
		--- PASS: TestCreateChannel (0.00s)
		PASS
		ok      membuat-channel 0.357s

		jadi kalian bisa close channelnya menggunakan defer
		atau langsung saja memanggil function close ya

		kalo misalkan data dari channel itu tidak kita gunakan atau kita terima
		maka akan terjadi error panic: send on closed channel

		jadi artinya kita mengirim data ke channel yang telah diclose
		dan kalo aklian perhatikan print yang ada didalam anonymous function
		tidak dijalankan

		sebaliknya kalo kita tidak mengirimkan data kechannel nya
		maka akan terjadi error all goroutines are sleep -  deadlock

		jadi pengirim dan penerima harus ada ya kalo salah satu tidak ada
		maka akan terjadi error

		jadi ini adalah komunikasi antar goroutine yaitu menggunakan channel
		itu lah cara membuat channel
		mudah mudahan kalian paham



	*/

}
