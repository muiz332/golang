/*

- sync map


golang memiliki sebuah struct dengan nama sync.map
sync map ini mirip dengan map biasa tetapi sync map ini aman digunakan untuk concurrency
artinya dia aman terhadap race condition

dan yang membedakan lagi sync.Map key dan valuenya itu bisa berbagai type data
kalo map biasa kan type data dari key dan valuenya kalo string harus string semua

nah kalo sync.Map ini kita tidak perlu menentukan type data key dan valuenya
artinnya kalian bisa bebas memasukkan type data apapun ya


ada beberapa method yang ada di sync.map

store => untuk menyimpan data ke map
load => untuk mengambil data ke map
delete => untuk menghapus data dari map
rage => untuk melakukan looping pada map

oke langsung saja kita coba ya


*/

package main

import (
	"fmt"
	"sync"
)

func addData(valueData int, data *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	data.Store(fmt.Sprintf("looping%d", valueData), valueData)
}

func main() {

	var data = &sync.Map{}
	var wg = &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go addData(i, data, wg)
	}

	wg.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	/*

		oke jadi seperti itu ya penggunakan dari sync.Map
		sangat sederhana sekali

		pada method range kita memasukkan sebuah function atau bisa dinamakan
		anonymous function yang mereturn boolean

		kenapa harus mereturn boolean?
		agar bisa menghentikan atau melanjutkan loopingnya

		kalo kalian return true maka loopingnya akan lanjut
		dan kalo kalian return false maka loopingnya akan terhenti

		jadi kalian tidak perlu menggunakan mutex lagi untuk menangani race condition
		mudah mudahan kalian paham

	*/

}
