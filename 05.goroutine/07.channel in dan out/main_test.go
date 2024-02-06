/*


- channel in dan out

saat kita mengirim channel sebagai parameter, function tersebut bisa menerima dan mengirim
data dari channel tersebut

kadang kita ingin memberitahu function tersebut akan hanya dapat menerima atau mengirim
data saja dari channel

nah kita bisa lakukan hal ini dengan menandai function tersebut sebagai in atau out
kalo kita atur sebagai in maka function tersebut hany bisa mengirim data dari channel

kalo kita atur sebagai out maka function tersebut hanya bisa menerima data dari channel
kalo kalian paksa untuk mengirim data kedalam channel maka akan terjadi error

oke langsung saja kita coba ya


*/

package main

import (
	"fmt"
	"testing"
)

func channelOnlyIn(channel chan<- string) {
	channel <- "ini data dari function channel only in"
}

func channelOnlyOut(channel <-chan string) {
	var data string = <-channel
	fmt.Println("function channel only out get data", data)
}

func TestChannelInOut(t *testing.T) {
	var channel chan string = make(chan string)
	defer close(channel)

	go channelOnlyIn(channel)
	go channelOnlyOut(channel)

	fmt.Println(<-channel)

	fmt.Println("program selesai")

	/*

		kalo kita jalankan hasilnya seperti ini ya

		ini data dari function channel only in
		program selesai
		--- PASS: TestChannelInOut (0.00s)
		function channel only out get data
		PASS
		ok      channel-in-out  0.362s


		jadi channel only outnya akan dijalankan secara bergantian atau concurrency

		jadi kalian bisa mengatur sebuah funciton itu hanya mengirim atau menerima data
		dari sebuah channel

		jadi seperti itu ya
		mudah mudahan kalian paham



	*/
}
