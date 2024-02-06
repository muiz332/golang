/*

- channel sebagai parameter

dalam pembuatan applikasi kita sering mengirim channel kedalam sebuah function
menggunakan parameter

sebelumnya kita tahu bahwa digolang defaultnya dari parameter itu adalah pas by value
artinya dicopy nilainya

nah khusus untuk channel secara default adalah pas by reference, jadi kalian tidak perlu
menggunakan pointer lagi untuk menjadikan pas by reference

oke langsung saja kita coba






*/

package main

import (
	"fmt"
	"testing"
	"time"
)

func getDataFromDatabase(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "ini data dari database"
}

func TestChannelAsParameter(t *testing.T) {

	var channel chan string = make(chan string)
	defer close(channel)
	go getDataFromDatabase(channel)

	fmt.Println(<-channel)
	fmt.Println("selesai dijalankan")

	/*

		jadi seperti itu ya cara untuk membuat channel sebagai parameter
		tinggal kalian gunakan keyword chan diikuti dengan type data dari
		channelnya

		oke mudah mudahan kalian paham

	*/

}
