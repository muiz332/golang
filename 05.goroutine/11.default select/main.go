/*

- default select


apa yang terjadi jika kita melakukan select pada channel yang tidak ada datanya?
maka akan terjadi error deadlock

maka kita harus menunggu sampai datanya ada, kadang kita ingin melakukan sesuatu selama proses
menunggu datanya tersebut

kalo kalian belajar javascript, dijavascript itu ada fitur finally,
jadi pada saat menunggu datanya kita bisa melakukan sesuai

digolang itu dinamakan default select

oke langsung saja ya kita coba



*/

package main

import (
	"fmt"
	"time"
)

func testing1(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "hello"
}
func testing2(channel chan<- string) {
	channel <- "hello kuy"
}

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go testing1(channel1)
	go testing2(channel2)

	i := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println(data)
			i++
		case data := <-channel2:
			fmt.Println(data)
			i++
		default:
			fmt.Println("sedang menunggu data")
		}

		if i == 2 {
			break
		}
	}

	/*

		jadi function testing1 itu saya berikan jeda dulu selama 1 detik,
		dari 1 detik itu kan tidak akan yang dilakukan

		maka default akan dijalankan selama proses menunggu tersebut

		jadi disini kalian tinggal menambahkan default saja didalam selectnya
		nanti default itu akan dijalankan selama channel belum menerima data

		atau dalam masa pending,
		jadi sesimple itu ya

		mudah muahan kalian paham

	*/

}
