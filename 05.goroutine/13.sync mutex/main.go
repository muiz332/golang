/*

- sync mutex

mutex itu singkatan dari mutual exclusion

untuk mengatasi masalah rece condition tersebut digolang terdapat sebuah struct bernama
sync.Mutex

mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking
terhadap mutex maka tidak ada yang bisa melakukan locking sampai kita melakukan unlock

jadi ketika nanti kalian memiliki satu buah variable yang disharing ke beberapa goroutine
sebelum kalian merubah variable tersebut kalian harus melakukan locking dan ketika melakukan
locking

si mutex ini hanya memperbolehkan satu goroutine yang melakukan locking,

jadi bayangin kalian memiliki beberapa goroutine, nantika ketika goroutine tersebut
ada yang berjalan secara paraller maka si mutex hanya memperbolehkan satu goroutine saja

yang dapat melakukan locking atau menjalankan satu goroutine yang dipilih secara random
dari beberapa goroutine yang berjalan secara paraller tersebut

jadi yang tadinya menjadi paraller akan berubah menjadi concurrency atau bergantian antar goroutine
setelah kalian unlock maka gorotuine selanjutnya akan dijalankan dan melakukan locking lagi

maka goroutine yang lain akan menunggu setelah goroutine tersebut melakukan unlock
ini sangat cocok sebagai solusi ketika ada masalah race condition
yang sebelumnya kita hadapi

oke langsung saja kita coba ya


*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var x = 0

	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {

			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)

	/*

		jadi sebelum melakukan perubahan terhadap variablenya kalian harus melakukan unclock
		terlebih dahulu

		setelah melakukan perubahan kalian harus unlock agar goroutine selanjutnya bisa dijalankan
		maka kalo misalkan ada goroutine yang berjalan secara paraller

		golang akan memiliki secara acak dari goroutine tersebut dan seakan akan membuat antrian
		antar goroutine

		kalo kita jalankan maka hasilnya akan benarnya ya yaitu 100000

		jadi kalo misalkan ada 1000 goroutine yang berjalan dan setelah satu goroutine
		itu berhasil melakukan lock

		maka goroutine yang lain akan menunggu goroutine tersebut ke unlock
		jadi itulah cara untuk mengatasi race condition


		lalu kapan kalian menggunakan mutex?
		kalo kalian memiliki kasus dimana sebuah variable itu disharing pada beberapa
		goroutine

		oke mudah mudahan kalian paham

	*/

}
