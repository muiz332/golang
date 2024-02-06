package test

import (
	"fmt"
	"menggagalkan-unit-test/helper"
	"testing"
)

func TestSayHelloMuiz(t *testing.T) {
	var result string = helper.SayHello("muiz")

	if result != "hello muiz" {
		t.Fail()
	}

	fmt.Println("selesai dijalankan muiz")
}

func TestSayHelloHasan(t *testing.T) {
	var result string = helper.SayHello("hasan")

	if result != "hello hasan" {
		t.FailNow()
	}

	fmt.Println("selesai dijalankan hasan")
}

func TestSayHelloHusain(t *testing.T) {
	var result string = helper.SayHello("husain")

	if result != "hello husain" {
		t.Error("harus mengembalikan hello husain")
	}

	fmt.Println("selesai dijalankan husain")
}

func TestSayHelloEko(t *testing.T) {
	var result string = helper.SayHello("eko")

	if result != "hello eko" {
		t.Error("harus mengembalikan hello eko")
	}

	fmt.Println("selesai dijalankan eko")
}

/*

kita jalankan menggunakan perintah

go test -v ./test

=== RUN   TestSayHelloMuiz
selesai dijalankan muiz
--- FAIL: TestSayHelloMuiz (0.00s)
=== RUN   TestSayHelloHasan
--- FAIL: TestSayHelloHasan (0.00s)
=== RUN   TestSayHelloHusain
    say_test.go:33: harus mengembalikan hello husain selesai dijalankan husain
--- FAIL: TestSayHelloHusain (0.00s)
=== RUN   TestSayHelloEko
    say_test.go:43: harus mengembalikan hello eko selesai dijalankan eko
--- FAIL: TestSayHelloEko (0.00s)


jadi simplenya kalo Fail() itu ketika ketemu fail maka code program yang dibawahnya
tetap bisa dijalankan, makanya muncul println selesai dijalankan muiz

kalo FailNow() itu code program yang dibawahnya tidka akan dijalankan
kalo Error() itu kita bisa menambahkan keterangan errornya dan dia akan menjalankan Fail()
artinya code program dibawahnya akan tetap dijalankan

kalo Fatal() itu kita juga bisa menambahkan keterangan errornya dan dia akan menjalankan
FailNow()  jadi code program yang dibawahnya tidak akan dijalankan

untuk rekomendasinya kalian gunakan yang Error() atau Fatal()
mungkin lebih sering yang Fatal ya

jadi itu lah cara untuk menggagalkan error
mudah mudahan kalian paham

*/
