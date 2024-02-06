/*


- membuat goroutine

untuk membuat goroutine digolang itu sangat sederhana
kita cukup menambahkan kata kunci go sebelum memanggil sebuah function yang akan kita jalankan
menggunakan goroutine

jadi kalo dijavascriptkan kita harus bikin promise lalu async await dulu
nah digolang kita cukup menggunakan kata kunci go saja

dan secara otomatis dia akan berjalan secara asyncronous menggunakan goroutine artinya tidak akan ditunggu
sampai function tersebut selesai

applikasi akan lanjut berjalan ke program yang selanjutnya tanpa menunggu goroutine yang kita buat
selesai

jadi hati hati kalo kalian tiba tiba menjalankan goroutinenya lalu tiba tiba applikasinya
berhenti maka goroutine tersebut tidak akan sampai dijalankan sampai selesai
karena sudah keburu berhenti

kita coba dengan menggunakan unit test ya


*/

package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello world")
}

func TestRunHelloWorld(t *testing.T) {

	go RunHelloWorld()
	fmt.Println("selamat datang")
}

/*

kalo kita jalankan hasilnya akan seperti ini

selamat datang
hello world

selamat datang
hello world

selamat datang

dari ketiga hasil tersebut kita bisa simpulkan bahwa ketika kita menambahkan goroutine pada function RunHelloWorld

maka function tersebut akan diskip sambil menjalankan code program yang ada dibawahnya
setelah selesai maka dia akan menjalankan goroutinenya lagi

maka hasilnya akan tampi selamat datang terlebih dahulu dan diikuti dengan hello world

akan tetapi ada juga yang hanya muncul selamat datang saja
hal ini terjadi karena code program kita sudah terlebih dahulu berhenti

jadikan setelah menjalankan print selamat datang program yang ada digoroutine tersebut
akan terhenti karena programnya sudah berhenti tidak bisa kembali ke goroutinenya lagi

karena itu kita bisa pending dulu programnya


*/

func TestRunHelloWorld2(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("selamat datang")

	time.Sleep(1 * time.Second)
}

/*

jadi kalian bisa tambahkan time.Sleep agar program bisa menunggu sebentar selama 1 detik sebelum
program benar benar berhenti

jadi itulah cara untuk membuat goroutine digolang
jadi sangat simple ya, kalian tinggal tambahkan kata kunci go agar sebuah function bisa berjalan secara
concurrency atau bergantian

mudah mudahan kalian paham



*/
