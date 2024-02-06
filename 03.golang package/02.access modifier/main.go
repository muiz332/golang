/*


- access modifier


dimateri kali ini kita akan belajar mengenai access modifier pada package

apa itu access modifier ?
access modifer adalah sebuah hak akses apakah sebuah code program yang berada didalam package tersebut
bisa kita aksess dari luar

simplenya kalo kalian membuat nama dari function, variable, struct dan lain lain
menggunakan awalan huruf besar, maka function tersebut bisa kita aksess dari luar package

sebalikanya, kalo kalian menggunakan awalan huruf kecil, maka function tersebut tidak akan
bisa diaksess dari luar package

oke langsung saja kita coba

jangan lupa untuk menuliskan
go mod ini nama-project

*/

package main

import (
	"access-modifier/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.Nama)
	utils.SayHello()

	/*

		kalo kalian paksa untuk mengakses function sayHi
		dimana function tersebut diawali dengan huruf kecil

		maka itu akan terjadi error

	*/

	// utils.sayHi()

	/*

		katanya
		undefined: utils.sayHi

		sayHi itu tidak ada, karena tadi, awalan katanya diawali dengan huruf kecil
		makanya ketika kita panggil diluar package utils tidak akan bisa terbaca

		jadi seperti itu ya materi kita tentang access modifier
		mudah mudahan kalian paham

	*/
}
