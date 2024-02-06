/*

- package

dimateri kali ini kita akan belajar mengenai  package
sebenarnya secara tidak langsugng kita sudah sering menggunakan package

yaitu ketika kalian menuliskan package main, itu menandakan bahwa difile inilah
program pertama kali dijalankan

jadi meskipun kalian memiliki banyak file digolang, akan tetapi hanya satu file
yang akan dijalankan

ketika program kalian sudah banyak, misalkan kalian memiliki 100 baris code
dengan menggunakan package kalian bisa memecah kompleksitas code program kalian

yang awalnya 100 menjadi beberapa folder yang terstruktur beserta file didalamnya
jadi simplenya package itu untuk merapikan code kita agar lebih mudah dibaca
dan juga memisahkan code agar kita tau, code mana yang menjalankan bagian A
dan code mana yang menjalankan bagian B atau istilahnya membuat grub

sederhananya package itu adalah sebuah folder yang sama dengan yang ada didalam
sistem operasi kita

sebenarnya ada 3 type macam package
yaitu package yang bisa kita buat sendiri atau build in package
yaitu package yang sudah disediakan oleh golang
dan terakhir third party package, atau package yang bisa kita install dan itu dibuat
oleh  orang lain

contoh package yang sudah disediakan oleh golang adalah package fmt
yang biasanya kita gunakan untuk menampilkan sesuatu diterminal

akan tetapi sekarang kita akan belajar membuat package sendiri
untuk membuat package kita harus menggunakan fitur yang namanya go modules


apa itu go modules ?
go modules merupakan manajemen package resmi dari go, go modulus diperkenalkan digolang versi go1.11


jadi ketika code kalian lebih complex kalian pasti menggunakan package untuk mengatur code kalian
dan tentunya kalian harus menggunakan fitur go modules agar package bisa digunakan

sekarang kita coba untuk menggunakan fitur go modules

jadi ketika kalian membuat project kalian bisa ketikkan

go mod init nama-project-kalian

untuk nama project sebenarnya direkomendasikan sesuai dengan nama folder kalian ya
cukup dengan mengetikkan seperti itu kalian sudah bisa menggunakan fitur package digolang
kalo kalian perhatikan akan ada sebuah file dengan nama go.mod

file itu berisi informasi informasi seperti package apa yang kita install
juga informasi mengenai go versi berapa yang sedang kita gunakan

sekarang kita coba fitur package digolang

*/

// package main

// import "fmt"

// func volumeKubus(s int) int {
// 	return s * s * s
// }

// func kelilingKubus(s int) int {
// 	return 12 * s
// }

// func main() {
// 	fmt.Println()
// 	fmt.Println("sebuah kubus memiliki panjang sisi 10 cm, hitunglah")
// 	fmt.Println("1.volume kubus\n2.keliling kubus")
// 	fmt.Println()

// 	var volume int = volumeKubus(10)
// 	var keliling int = kelilingKubus(10)
// 	fmt.Println("volume kubus adalah", volume, "cm")
// 	fmt.Println("keliling kubus ", keliling, "cm")
// 	fmt.Println()

// }

/*

	misalkan kita memiliki 2 buah function seperti itu
	bayangkan ketika kalian nanti sudah mengerjakan program yang lebih besar

	tentu kalian punya banyak function function, itu akan membuat code program kalian
	menumpuk dan tidak rapi

	karena berbagai macam function yang mengerjakan hal yang berbeda
	ada didalam satu file

	nah kita bisa gunakan fitur yang namanya package untuk merapikan code atau
	membuat grub code kita

	yang pertama kalian buat foldernya dulu dengan nama kubus
	dan buat file dengan nama kubus.go

	nah kita bisa pindahkan function yang berhubungan dengan kubus
	didalam file tersebut

	dengan catatan nama function harus diawali dengan huruf besar
	agar kita bisa menggunakan function tersebut diluar package kubus

	maka kita bisa menggunakan yang namanya import
	jadi simplenya import itu memanggil code program kita yang berada
	didalam package seperti ini


	import (
		"belajar-package/kubus"
		"fmt"
	)

	dapat dari mana belajar-package itu?
	itu adalah nama project kita yang kita tuliskan saat
	menjalankan go mod init

	akan tetapi kalian bisa memanggil nama packagenya
	diikuti dengan nama functionnya

	secara otomatis akan auto import


*/

package main

import (
	"belajar-package/kubus"
	"fmt"
)

func main() {

	fmt.Println()
	fmt.Println("sebuah kubus memiliki panjang sisi 10 cm, hitunglah")
	fmt.Println("1.volume kubus\n2.keliling kubus")
	fmt.Println()

	var volume int = kubus.VolumeKubus(10)
	var keliling int = kubus.KelilingKubus(10)
	fmt.Println("volume kubus adalah", volume, "cm")
	fmt.Println("keliling kubus ", keliling, "cm")
	fmt.Println()

	/*

		kalo kita jalanakan maka akan muncul hasilnya
		jadi seperti itu lah fitur package digolang

		itu akan mempermudah untuk mengelompokkan code program
		kita, agar code progam kita lebih rapi dan mudah untuk dibaca

		kalian tidak bisa membuat nama yang sama dalam satu package
		meskipun itu berbeda file

		tapi kalo packagenya berbeda itu bisa ya
		oke jadi seperti itu materi tentang package
		mudah mudahan kalian paham

	*/
}
