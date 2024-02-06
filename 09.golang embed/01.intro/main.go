/*

- golang embed

sejak golng versi 1.16 terdapat package baru dengan nama embed
package embed adalah fitur baru untuk mempermudah membaca isi file pada saat
melakukan compile

dan secara otomatis file external yang kita baca akan disertakan menjadi satu file
pada program golang saat mengcompile code program golang

jadi antara file external dan program golang kita, ketika proses compilasi berlangsung
maka file external dan program golang kita akan disatukan menjadi satu buah file
golang yang telah dicompile

jadi nanti misalkan ada sebuah file.txt kita lakukan embed
maka kita bisa baca file tersebut menjadi file biner dan ikut dalam hasil compilenya

nah cara menambah embed kalian bisa tambahkan komentar

//go:embed diikuti dengan nama file yang mau kita baca dan komentar ini kita tulis diatas variable

maka secara otomatis variable tersebut akan berisi data sesuai dengan isifilenya
tetapi sebelum itu kalian jangan lupa untuk melakukan import embed ya dengan memanggil funciton
initnya saja

artinya kalian tambahkan underscore didepan pemanggilan packagenya
jadi nanti ketika dicompile isi filenya akan masuk kedalam isi variablenya

dan kalian tidak bisa bikin golang embed didalam function
jadi variablenya tidak bisa didalam function

jadi seperti itu ya simplenya embed ini
mudah mudahan kalian paha



*/

package main

func main() {

}
