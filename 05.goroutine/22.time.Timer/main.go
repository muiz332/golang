/*

- timer.Timer

timer adalah representasi dari satu kejadian
jadi simplenya kita bisa lakukan delay untuk menjalankan code program kita

kalo kalian belajar javascript, itu ada fitur yang namanya setTimeout
nah sama seperti itu

tetapi digolang menggunakan timer
ketika waktu timer itu sudah habis dari waktu yang telah ditentukan
maka event akan dikirimkan ke channel

untuk membuat timer.Timer kita bisa gunakan
time.NewTimer(duration)

oke langsung saja kita coba



*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func TimerAfter() {
	tunggu := time.After(1 * time.Second)
	<-tunggu
	fmt.Println("selesai timer after ")
}

func AfterFunc(wg *sync.WaitGroup) {

	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		defer wg.Done()
		fmt.Println("ini after func selesai")
	})
}

func main() {

	var timer = time.NewTimer(5 * time.Second)
	fmt.Println(time.Now().Second())

	time := <-timer.C // menunggu waktu habis
	fmt.Println(time.Second())
	fmt.Println("selesai dijalankan")

	/*

		kalo kita jalankan maka dia akan menunggu selama 5 detik barulah
		print selesai dijalankan akan tampil

		lalu ada time.After()
		jadi ini sama saja tetapi data yang dikembalikan langsung berupa channel ya
		kita coba

	*/

	TimerAfter()

	/*

		nah selanjutnya ada afterFunc, nah ini kalo kalian ingin menunggu dulu lalu ingin
		menjalanka sesuatu didalam sebuah function

		afterfunc akan berjalan secara asyncronous menggunakan goroutine, maka dari itu
		kalo kita ingin menunggu kita bisa gunakan waitGroup

		oke langsung saja kita coba

	*/

	var wg sync.WaitGroup
	AfterFunc(&wg)
	wg.Wait()
	fmt.Println("semua program selesai dijalankan")

	/*

		jadi itulah mengenai timer.Timer

		mudah mudahan kalian paham

	*/

}
