package test

import (
	"membuat-unit-test/helper"
	"testing"
)

func TestSayHelloName(t *testing.T) {

	var result string = helper.SayHelloName("mui")

	if result != "Hello muiz" {
		panic("Result is not 'Hello muiz'")
	}
}

/*

jadi kalian buatnya seperti itu ya, menggunakan nama awalan Test difunction unit testnya
dan apa yang dilakukan difunction tersebut ?

tentu saja kita harus mengetest, misalkan function SayHelloName dengan parameter
muiz, itu harus mengembalikan Hello muiz

kalo tidak sama dengan Hello muiz maka saya kasih panic
dan sangat simple

kalian cukup membuat sekenario testingnya, dan yang kalian harapkan hasilnya
dari sekenario tersebut itu apa

dan kita coba jalankan

untuk menjalankannya kalian cukup tuliskan

go test => menjalankan semua unit testnya
go test -v = > menjalankan semua unit testnya dan menampilkan nama unit testnya, v untuk verbose
go test -v -run TestNamaFunction => menjalankan unit test tertentu

dengan catatan kalian harus masuk kedalam foldernya atau packagenya

atau kalo kalian menggunakan vs code kalian cukup tekan tombol run disebelahh
kiri function unit testnya dan ini kita bisa jalankan diluar package test yang kita buat

maka akan muncul seperti ini kalo kalian menggunakan -v atau menggunakan tombol run

=== RUN   TestSayHelloName
--- PASS: TestSayHelloName (0.00s)
PASS
ok      membuat-unit-test/test  0.235s

lalu bagaimana kalo saya ingin menjalankan diluar package test yang kita buat
atau di root folder kita ?

kita gunakan perintah

go test -v ./...

jadi dia akan menjalankan semua unit testnya
jadi seperti itu ya unit test digolang

mudah mudahan kalian paham

*/
