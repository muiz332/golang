/*

- select channel


kadang ada kasus dimana kita membuat beberapa channel dan menjalankan bebarapa goroutine
lalu kita ingin mendapatkan data dari semua channel tersebut

jadi misalkan kita punya 2 buah channel dan 2 buah channel tersebut kita tidak tahu
berapa banyak data yang dikirimkan dichannelnya

nah kalo kita menggunakan for range itu hanya untuk satu channel, untuk mengatasi hal tersebut
kita menggunakan select channel

dengan select channel kita bisa memilih data tercepat dari bebarapa channel, jika data datang secara
bersamaan dibeberapa channel, maka akan dipilih secara random

jadi nanti kalian tinggal sebutin nama nama channelnya saja


jadi sebenarnya select itu hanya mengambil data satu kali pada sebuah channel
jadi nanti kita hanya perlu menggunakan for yang didalamnya terdapat select

langsung saja kita coba ya






*/

package main

import (
	"fmt"
)

func testing1(channel chan<- string) {
	defer close(channel)
	channel <- "testing1"
}
func testing2(channel chan<- string) {
	defer close(channel)
	channel <- "testing2"
}

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go testing1(channel1)
	go testing2(channel2)

	var i int8 = 0

	for {
		select {
		case data := <-channel1:
			fmt.Println(data)
			i++
		case data := <-channel2:
			fmt.Println(data)
			i++
		}

		if i == 2 {
			break
		}
	}

	/*

		jadi seperti itu ya cara penggunaan dari select channel
		kalian bisa menangkap data dari beberapa channel dimulai dari data yang paling cepat
		dikirimkan

		kalo misalkan kalian tidak menggunakan select, maka yang terjadi adalah saling menunggu
		atau blocking

		jadi channel2 harus menunggu dulu sampai channel ke 1 selesai
		barulah channel ke2 dijalankan

		jadi seperti itu ya select channel mudah mudahan kalian paham

	*/

}
