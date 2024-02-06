/*

- golang unit test


dimateri kali ini kita akan bahas mengenai golang unit test

sebelum kalian membuat applikasi yang baik dan benar, kalian harus belajar dulu
mengenai unit test


secara umum ada 3 jenis unit test


UI
service
unit


jadi UI test itu akan melakukan testing dari sisi usernya, jadi setelah applikasinya jadi
kita test dari sisi usernya, kalo applikasi web berarti dari sisi UInya

lalu ada service test
jadi didalam applikasi kita ada applikasi yang lain

misalkan 1 front end memiliki 2 buah backend, maka service test itu akan mengetest tiap tiap backend tersebut

lalu ada  unit test, didalam applikasi miisalkan backend, itu ada modul module atau function function
tertetu, nah unit test itu melakukan pengujian terhadap funciton function yang kita buat

lalu kenapa kita test?
supaya code program kita minim kesalahan atau bug

jadi nanti kalian mengetestnya harus memiliki sekenario yang banyak ya, mulai dari sekenario berhasil dan gagal
ketika berhasil apa yang dilakukan dan ketika gagal apa yang dilakukan

jadi misalkan kalian memiliki 1 buah funciton, bisa saja nanti unit test yang kalian buat itu ada 4
jadi akan lebih banyak unit testnya daripada codenya ya


digolang untuk melakukan pengetest an itu sangat mudah, berbeda dengan bahasa pemrograman yang lain
seperti javacript yang mengharuskan kita menginstall framework jest sebagai softwere testingnya

digolang itu sudah disediakan sebuah package yang namanya testing yang dapat kita gunakan
untuk membuat unit test

jadi tinggal kalian pakai saja ya
kita akan coba dimateri berikutnya

mudah mudahan kalian bisa mengikuti materi ini sampai akhir


*/

package main

import "fmt"

func main() {
	fmt.Println("Golang unit test")
}
