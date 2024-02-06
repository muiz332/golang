package main

import "fmt"

func main() {
	fmt.Println("Hello World")

}

/*

jadi kita harus membuat func main yang didalamnya adalah program kita
kita akan menuliskan program didalam func main tersebut

untuk menampilkannya ke terminal kalian tuliskan
fmt.Println() yang didalamnya terdapat tulisan hello world


fmt adalah package bawaanya go, jadi package itu adalah sebuah program
yang bisa langsung kita pakai dan program tersebut sudah disediakan
oleh golang

jadi kita tinggal pakai saja


selanjutnya untuk melakukan kompilasi kalian cukup tuliskan

go build namafile.go

kalo kita jalankan maka akan akan file dengan nama
main.exe kalo diwindows, jadi tergantung sistem operasi apa yang
kalian gunakan

dan untuk menjalankannya kalian tuliskan nama filenya saja
diterminal kalian tuliskan

./main

maka akan muncul Hello World



akan tetapi agak ribet ya kalo kalian bikin file go
kemudian ketika ingin kita jalankan harus mengcompile dulu

untungnya digolang itu kita disediakan perintah yang bisa menjalankan
program kita tanpa harus kita compile

akan tetapi ini tersedia difase development saja ya
setelah program selesai, kita harus compile dulu

kalian tulis

go run namafile.go

maka akan muncul tulisan hello world






*/
