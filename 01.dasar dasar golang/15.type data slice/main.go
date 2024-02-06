/*

- type data slice


kali ini kita akan bahas mengenai type data digolang
slice adalah type data yang berbeda pada dengan type data yang lainnya


apa itu type data slice?
type data slice adalah type data potongan dari array

jadi slice itu mirip dengan array, cara mengaksesnya, cara mengubahnya, dia punya index
jadi dia mirip dengan array

cuma yang membedakan ukuran slice itu bisa beruba, kalo array itukan ketika kita sudah
tentukan panjang arraynya ya sudah tidak bisa bertambah lagi

kalo didalam slice itu bisa, kalian bisa mengambah data didalam slice
jadi mirip dynamic ya, tapi tidak dynamic juga

slice dan array itu selalu terkoneksi, dimana slice itu adalah data
yang mengakses sebagian atau seluruh data didalam array

jadi sebenarnya ketika kalian nanti bikin slice, sebenarnya dibelakan slice itu ada
array, jadi tempat menyimpan data itu sebenarnya array juga

cuma didepannya itu ada slice, kita berkomunikasi menggunakan slice, akan tetapi
datanya tetap disimpan didalam array

type data slice itu memiliki 3 data,
ada pointer, length, dan capacity

pointer adalah penunjuk ada pertama diarray pada slice
length adalah panjang dari slice
capacity adalah kapasitas atau daya tampung data slicenya, dimana length tidak boleh
lebih dari capacity

jadi kalo capacitynya itu 10 maka lengthnya itu 10 atau dibawah 10


untuk membuat array kalian bisa tulis seperti ini

array[low:high]    => membuat slice dari array low sampai index sebelum high
array[low:]		  => membuat slice dari low sampai total data array
array[:hight]	  => membuat data slice dari index 0 sampai sebelum high
array[:]		  => membuat dara slice dari index 0 sampai total data array

jadi slice itu potongan dari array
kalian bisa memotong array seperti itu

oke langsung saja kita coba





*/

package main

import "fmt"

func main() {

	var bulan [12]string = [12]string{
		"january",
		"february",
		"maret",
		"april",
		"mei",
		"juni",
		"july",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 []string = bulan[4:7]
	var slice2 []string = bulan[6:9]
	fmt.Println(slice1) // [mei juni july]
	fmt.Println(slice2) // [july agustus september]
	fmt.Println()

	/*

		kalo kita jalankan maka hasilnya seperti ini [mei juni july]
		inilah yang disebut dengan slice yaitu hasil potongan dari array

		pointernya adalah 4
		lengthnya adalah 3
		capacitynya adalah 8

		index ke 0 dislicenya kan mei, mei index berapa diarray yang kita potong?
		yaitu index ke 4, index 4 itulah pointernya

		capacitynya itu dari array index ke 4 sampai index ke akhir ada berapa?
		ada 8 data, itulah capacity didalam slicenya

		jadi slicenya hanya bisa menampung total 8 data
		kalo lebih dia akan membuat array baru

		sekarang apa yang akan terjadi kalo saya mengubah index ke 0 dislicenya

	*/

	slice1[0] = "meikuy"
	fmt.Println(slice1)
	fmt.Println(bulan)
	fmt.Println()

	/*

		maka slicenya akan berubah, tetapi lihat array yang kita potong, meinya juga ikut berubah menjadi meikuy

		kenapa seperti ini? ini karena ketika kita memotong arraynya menjadi slice, maka slice itu
		akan tetap mengacu kepada array yang dipotong

		jadi istilahnya slicenya mereference array yang dipotong dalan hal ini array bulan
		begitu pula sebaliknya

		kalo kalian mengubah arraynya maka slice yang mengacu kearray tersebut akan ikut berubah
		juga
	*/

	bulan[6] = "julykuy"
	fmt.Println(slice1)
	fmt.Println(slice2)

	/*
		maka slice1 dan slice2 akan ikut berubah juga
		jadi hati hati kalo kalian ingin mengubah slice atau array

	*/

}
