/*

- ticker

ticker ada representasi kejadian berulang

jadi kalian bisa membuat kejadian yang berulang ulang setiap berapa waktu
untuk membuat ticker kalian bisa tulis seperti ini

time.NewTicker(duration)
dan untuk menghentikan ticker kalian bisa tulis ticker.Stop()

jadi ini sama dengan setInterval pada javascript

oke langsung saja kita coba


*/

package main

import (
	"fmt"
	"time"
)

func tickerSelect() {
	done := make(chan bool)
	ulang := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
		return
	}()

	for {
		select {
		case <-done:
			ulang.Stop()
			return
		case t := <-ulang.C:
			fmt.Println("Hello !!", t)
		}
	}
}

func main() {

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {

		fmt.Println(tick)
	}

	/*

		jadi ini golang memanfaatkan channel yang dikirim datanya secara terus menerus
		dengan cara mengirimkan data ke channelnya dengan delay waktu yang ditentukan
		pada parameter newTickernya

		jadi kalo sebuah channel itu terus menerus dikirimkan tanpa henti
		maka itu akan menciptakan program yang dijalankan tanpa henti dalam durasi waktu tertentu

		nah tinggal kita tangkap saja data dari channel yang telah dikirimkan oleh tickernya
		jadi kita tidak tahu berapa banyak data yang dikirimkan dichannelnya

		maka dari itu kita gunakan for range, harusnya data yang dikirimkan ke
		channel itu tidak akan ada habisnya ya

		mungkin dia menggunakan looping tanpa batas, untuk menghentikannya kalian tekan control c
		diterminal

		nah kita bisa hentikan, misalkan saya ingin hentikan dalam waktu 5 detik
		jadi seperti yang kalian lihat didalam anonymous functionnya

		tetapi akan terjadi error deadlock, karena tickernya sudah berhenti mengirim data ke channel
		kalo mau tidak error kalian bisa gunakan select seperti ini


	*/

	// tickerSelect()

	/*

		jadi kalian tinggal buat satu channel lagi untuk memberitahu tickernya berhenti atau tidak
		dalam kasus ini nama channelnya adalah done

		didalam select untuk menghentikan tickernya kalian tinggal tuliskan sleep
		kemudian diikuti dengan return ya

		karena kalo tidak kita return maka akan terjadi error deadlock
		lalu kenapa returnnya tidak didalam goroutine saja?

		jangan lupa goroutine itu tidak dapat melakukan return, jadi akan sama saja terjadi
		error deadlock

		jadi seperti itu ya ticker digolang
		mudah mudahan kalian paham

	*/
}
