/*

- range channel


kadang ada kasus dimana channel dikirim data secara terus menerus oleh pengirim
dan tidak jelas kapan channel tersebut akan berhenti menerima data

salah satu cara yang bisa kita lakukan adalah menggunakan perulangan range ketika
kita menerima data dari channel

ketika sebuah channel diclose maka secara otomatis perulangan tersebut akan terhenti
ini lebih sederhana dari pada kita melakukan pengecheckan secara manual

oke langsung saja kita coba ya


*/

package main

import (
	"fmt"
	"strconv"
)

func sendChannel(channel chan<- string) {
	defer close(channel)

	for i := 0; i < 10; i++ {
		fmt.Println("ke", i)
		channel <- "mengirim data ke " + strconv.Itoa(i)
	}
}

func main() {

	var channel chan string = make(chan string)

	go sendChannel(channel)

	for val := range channel {
		fmt.Println(string(val))
	}

	/*

		jadi misalkan kita punya sebuah function yang dibuat oleh orang lain
		namanya sendChannel dan kita tidak tahu channel yang dikirim

		nah misalkan difunction tersebut kita kirim 10 data ke channel

		maka setelah function tersebut mengirim maka akan langsung diterima oleh channelnya
		dan menunggu ada yang mengambil data tersebut

		karena kita tidka tahu berapa data yang dikirim maka
		kita gunakan perulangan untuk mengambil data dari channel tersebut

		jadi prosesnya secara concurrency ya bergantian
		dengan catatan kita harus menggunakan defer setelah selesai mengirim datanya
		pada function send channel

		jadi prosesnya secara bergantian

		mengirim data ke channel, menunggu ada yang menerima, lalu data sudah diterima
		mengirim data ke channel, menunggu ada yang menerima, lalu data sudah diterima

		seperti itu sampai 10 data yang dikirim pada channel karena pada function
		sendChannel itu mengirim data menggunakan looping sebanyak 10 kali

		oke jadi seperti itu ya range channel

		bagaimana kita menerima lebih dari satu data dari channel, tetapi kita tidak tahu
		berapa data yang dikirimkan didalam channel tersebut

		oke mudah mudahan kalian paham


	*/
}
