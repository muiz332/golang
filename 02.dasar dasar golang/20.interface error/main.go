/*


- interface error

dimateri kali ini kita akan membahas mengenai interface error

golang memilik interface yang bisa kita gunakan untuk mengimplemtasikan interface
error, nama interfacenya adalah error

type error interface {
	Error() string
}

untuk membuat error, kalian tidak perlu membuat struct yang mengimplementasikan interface
error ini

golang sudah membuat itu, jadi kita tinggal panggil saja package errors dengan
cara menuliskan

errors.New(msg)

apa itu package? nanti kita akan bahas ya
untuk sekarang kalian ikuti dulu saja

umumnya digolang ketika kita membuat sebuah function yang berkemungkinan
adanya error, maka kita bisa mengimplementasikan interface error ini

jadi nanti functionnya akan mengembalikan 2 multiple value
yang satu value yang asli dan yang kedua value yang memberitahu apakah function
tersebut error atau tidak


jadi kita coba bikin sebuah function pembagian
dimana kalo sebuah angka dibagi dengan 0

digolang akan terjadi panic, nah kita bisa check kalo
pembaginya 0 kita bisa return interface error ini

kita coba

*/

package main

import (
	"errors"
	"fmt"
)

func pembagian(angka, bagi int) (int, error) {
	if bagi == 0 {
		return 0, errors.New("angka tidak bisa dibagi dengan 0")
	}

	return angka / bagi, nil
}

func main() {

	var result int
	var err error

	result, err = pembagian(10, 0)
	if err == nil {
		fmt.Println("hasilnya ", result)
	} else {
		fmt.Println("terjadi kesalahan", err.Error())
	}

	/*

		nah kalo saya jalankan, saya masukkan parameter pertama itu 10 dan parameter
		kedua itu 0

		maka didalam function pembagian akan dicheck
		apakah bagi itu sama dengan 0 ? true

		kalo true kita return 0, package error
		kalo false kembalikan hasilnya dan nil

		ingat ya nil ini adalah sebuah value yang menandakan bahwa itu adalah data kosong
		jadi karena error ini adalah interface kalo tidak error kita bisa return nil


		ketika function pembagian itu mengembalikan nilainya result dan err
		kita bisa check

		apakah err itu nilai nil ? false
		maka masuk ke else dan menampilkan kelayar pesan errornya

		dari variable err untuk mengambil pesan error kita bisa gunakan
		method Error()

		kalo saya jalankan hasilnya seperti ini

		terjadi kesalahan angka tidak bisa dibagi dengan 0

		artinya kita sudah berhasil membuat interface error

		jadi kalo kalian membuat sebuah function yang memiliki potensi terjadinya
		panic, maka kalian bisa implementasikan interface error ini

		jadi itulah interface error
		mudah muahan kalian paham

	*/

}
