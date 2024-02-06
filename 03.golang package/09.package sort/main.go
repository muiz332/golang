/*

- package sort


dimateri kali ini kita akan belajar mengenai package sort digolang
jadi package ini digunakan untuk mengurutkan slice ya

oke langsung saja kita coba

*/

package main

import (
	"fmt"
	"sort"
)

type Nilai []int

func (a Nilai) Len() int           { return len(a) }
func (a Nilai) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nilai) Less(i, j int) bool { return a[i] > a[j] }

func main() {

	// cara pertama

	var angka []int8 = []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sort.Slice(angka, func(i, j int) bool {
		return angka[i] > angka[j]
	})

	fmt.Println(angka)

	/*

		jadi kalian tinggal gunakan package sort.Slice
		dimana terdapat 2 buah parameter, parameter yang pertama itu adalah slicenya
		dan parameter kedua itu adalah function yang kita bisa gunakan untuk mengurutkannya

		jadi ketika kandisi dan direturn itu bernailai false
		maka akan dilakukan swap

		jadi sesimple itu ya untuk mengurutkan slice

	*/

	// cara kedua

	var nilai Nilai = Nilai{1, 2, 3, 4, 5, 6}
	sort.Sort(Nilai(nilai))

	fmt.Println(nilai)

	/*

		untuk cara kedua ini kalian harus definisikan dulu 3 buah function karena cara ini mengimplementasikan
		interface, ada funciton Len, Less, dan Swap

		dimana funciton Less digunakan untuk mengecheck kapan kita harus melakukan penukaran data
		dan function swap digunakan untuk menukar data

		nah disini kita bisa memanipulasi bagaimana cara mengecheck kapan kita mmelakukan penukaran data
		dan bagaimana data itu ditukar

		jadi cara kedua ini memiliki banyak control yang dapat kita gunakan
		akan tetapi ini jarang digunakan

		kalo kalian hanya mengurutkan dari yang terbesar keterkecil misalkan
		kalian cukup gunakan cara yang pertama ya

		oke jadi itu lah package sort
		mudah mudahan kalian paham

	*/

}
