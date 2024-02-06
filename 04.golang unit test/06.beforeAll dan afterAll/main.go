/*

- beforeAll dan AfterAll

biasanya didalam unit test kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test
dieksekusi

kita bisa menggunakan fitur testing.M
fitur ini bernama main, dimana digunakan untuk mengatur eksekusi unit test

sebenarnya digolang tidak ada fitur before dan after test, kita bisa gunakan fitur testing.M
ini agar seolah olah kita menggunakan before dan after

untuk mengatur eksekusi test kita cukup membuat function yang namanya TestMain
dengan parameter *testing.M


jika ada function TestMain maka secara otomatis golang akan mengeksekusi function ini tiap kali akan
menjalankan unit test disebuah package

dengan ini kita bisa mengatur before dan after pada unit test sesuai dengan yang
kita mau

dan TestMain ini akan dieksekusi sekali saja setiap kita menjalankan unit testnya
jadi sebelum semua unit testnya dijalankan makan function TestMain ini akan dijalankan dulu

oke langsung saja kita coba


*/

package main

func main() {

}
