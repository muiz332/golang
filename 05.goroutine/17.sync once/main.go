/*

- sync once

once adalah fitur digolang yang bisa kita gunakan untuk memastikan
sebuah function diexekusi hanya sekali pada goroutine

jadi berapa banyak goroutine yang mengakses function tersebut hanya dieksekusi
sekali saja

dan goroutine yang lain akan menghiraukannya
oke langsung saja kita coba






*/

package main

import (
	"fmt"
	"sync"
)

func sayHello() {
	fmt.Println("hello")
}

func main() {

	var group sync.WaitGroup
	var once sync.Once

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			once.Do(sayHello) // hanya dijalankan sekali saja
			fmt.Println("goroutine dijalankan")
		}()
	}

	group.Wait()
	fmt.Println("selesai")

	/*

		oke jadi seperti itu ya penggunaan dari sync.once
		mudah mudahan kalian paham

	*/

}
