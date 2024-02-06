/*

- struct

apa itu struct?
struct adalah sebuah template data yang dapat kita gunakan untuk menggabungkan satu atau lebih
type data yang berbeda

struct biasanya digunakan untuk merepresentasikan sebuah object
data struct disimpan kedalam field

sederhananya struct adalah kumpulan field

jadi kalo kalian pernah belajar bahasa pemrograman OOP
struct ini mirip dengan class

jadi kita bisa buat class orang yang didalamnya ada beberapa field
misalkan nama orang, umur dan lain lain

nah struct mirip seperti itu, jadi kita bisa buat representasi data


nah sebelumnya kan kita menggunakan array sebagai tempat menyimpan
data, tapi masalahnya dia hanya bisa diakses menggunakan index yang bertype integer

nah kalo map juga bisa digunakan, akan tetapi key nya harus memiliki type data yang
sama semua, jadi kalo typenya string ya string semua

dan juga valuenya, kalo valuenya bertype string harus string semua tidak bisa
macam macam type data, dan untuk aksesnya menggunakan index dengan type data
yang telah dituliskan

kalo struct ini kita bisa menyimpan value dengan type data yang berbeda beda
dan untuk aksesnya menggunakan nama fieldnya dimana kita bisa memberi
nama pada fieldnya

jadi kalo kalian dari bahasa pemrograman OOP ketika bikin class orang
sekarang kita bikin struct orang

kita coba ya

*/

package main

import "fmt"

type Orang struct {
	Nama, Alamat string
	Umur         uint8
}

func main() {

	// cara sederhana

	var hasan Orang
	hasan.Nama = "hasan"
	hasan.Alamat = "banyuwangi"
	hasan.Umur = 10

	fmt.Println(hasan)
	fmt.Println(hasan.Nama)
	fmt.Println(hasan.Alamat)
	fmt.Println(hasan.Umur)

	/*
		jadi untuk membuat struct, kita harus buat menggunakan keyword type
		diikuti dengan nama structnya, biasanya menggunakan huruf besar diawal, dan
		diikuti dengan keyword struct

		nah didalam type structnya kalian bisa memasukkan
		field yang dibutuhkan beserta type datanya

		jadi structny ini seperti membuat kumpulan type data baru ya
		setelah type data struct Orang dibuat

		maka kita gunakan, misalkan pada variable hasan dilanjutkan dengan type datanya yaitu
		struct Orang dan kita assigment dengan isi structnya

		misalkan field nama diisi dengan hasan, alamat disi dengan banyuwangi dan umur diisi dengan 10
		untuk field kalian gunakan awalan huruf besar ya, kalo lebih dari satu kata

		setiap kata harus diawali dengan huruf besar ya
		jadi seperti itu lah struct

		representasi dari kumpulan type data anggap saja kalian membuat object
		jadi sebenarnya struct ini adalah sebuah type data

		dimana type data struct ini tidak bisa langsung kita pakai
		kita harus mengisinya terlebih dahulu pada variable hasan

		setelah kita assigment, barulah kita bisa pakai, jadi struct juga bisa digunakan
		sebapai parameter dari sebuah function ya


		selanjutnya kita akan coba macam macam cara untuk membuat struct
	*/

	// cara paling direkomendasikan

	var muiz Orang = Orang{
		Nama:   "muiz",
		Alamat: "banyuwangi",
		Umur:   10,
	}
	fmt.Println(muiz)

	// cara singkat tetapi tidak direkomendasikan

	var husain Orang = Orang{"husain", "banyuwangi", 10}
	fmt.Println(husain)

	/*

		jadi untuk cara singkan kenapa tidak direkomendasikan?
		karena ketika type dari structnya itu ada penambahan field

		maka cara singkat akan terjadi error sedangkan cara yang lain
		tidak akan error

		lalu cara singkan itu tergantung pada urutan dari penempatannya
		kalo misalkan tertukar itu akan jadi masalah

		jadi karena itulah gunakan cara sederhana atau cara yang direkomendasikan
		mudah muahan kalian paham

	*/

}
