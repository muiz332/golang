/*

- membuat unit test


golang menyediakan sebuah struct yang bernama testing.T
struct ini digunakan untuk unit test digolang

jadi nanti ketika kalian membuat unit test, kalian akan menggunakan struct T ini

selanjutnya ada yang namanya testing.M, ini digunakan untuk mengatur life cycle testing

lalu ada juga testing.B, ini digunakan untuk benchmarking
untuk melakukan benhcmark (mengukur kecepatan kode program) itu sudah disediakan
sehingga kita tidak perlu menggunakan library atau framework tambahan



sekarang kita akan coba membuat unit testnya
tetapi sebelum itu kita contoh sebuah function, dan kita akan bikin unit test
dari function tersebut

dan kita simpan function tersebut kedalam sebuah package yang namanya helper


setelah kita bikin function yang akan ditest, selanjutnya kita buat unit testnya
ada beberapa aturan yang harus kalian perhatikan

untuk membuat file unit test kalian harus akhiri katanya dengan _test.go
jadi misalkan say_test.go

jadi kalian harus bikin nama file unit test itu sesuai dengan file mana yang mau kalian test
agar lebih mudah dibaca dan tidak membingungkan

dan aturan nama functionnya, digolang kalian harus menuliskan nama function unit test diawali
dengan kata Test dan kata selanjutnya bebas

dan selanjutnya harus memiliki parameter pointer T, *testing.T
dan dia tidak boleh mengembalikan return value

sekarang kita coba



*/

package main

func main() {

}
