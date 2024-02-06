/*

- Type data number

didalam golang, type data number itu ada dua jenis

- integer
- floating point

integer itu adalah sebuah angka desimal ya, dan floating point atau float, itu adalah angka
yang memiliki koma





- Integer

sekarang kita lihat macam macam type data integer

type data				nilai minimum 					nilai maximum
int8						-128							127
int16					   -32768					      -32767
int32					-2147483648						2147483647
int64				-9223372036854775808			9223372036854775807

jadi semakin besar kapasitas nya, semakin besar juga memory yang digunakan
jadi kalian silahkan gunakan type data yang sesuai dengan kebutuhan

misalkan umur, jadi nanti kalian gunakan int8 ya untuk umur
jangan menggunakan int16, karena umu orang rata rata tidak sampai 100 tahun

kalo kalian mau menyimpan saldo misalkan, kalian bisa gunakan int32 atau int64
jadi sesuai dengan kebutuhan kalian masing masing ya



type data				nilai minimum 					nilai maximum
uint8						0								225
uint16					    0					      	   65535
uint32						0							4294967295
uint64						0						18446744073709551615

ada juga uint, dimana dia tidak memiliki nilai negatif dan nilai maximumnya naik,
jadi seperti nilai minimum dalam bentuk positif ditambah dengan nilai maximum,

jadi nanti kalo kalian memiliki nilai yang sekiranya tidak mungkin untuk bernilai
negatif maka kalian bisa gunakan uint

contohnya umur, umur kalian bisa gunakan unit8, karena umur itu tidak bisa negatif



- Floating point


type data				nilai minimum 					nilai maximum
float32						1.18×10 pangkat 38		3.4×10 pangkat 38
float64					   2.23×10 pangkat 308		1.80×10 pangkat 308
complex32			complex numbers with float32 real and imaginary parts
complex64			complex numbers with float64 real and imaginary parts


jadi yang sering dipakai adalah yang float32 dan juga float64
yang complex itu jangan dipakai, kecuali kalo kita membangung
sebuah applikasi matematik yang membutuhkan perhitungan tingkat lanjut




- Alias

ada beberapa alias atau nama lain untuk type data yang tersedia


type data				alias untuk
byte						int8
rune						int32
int 		int tergantung sistem operasi 32 atau 64
uint		uint tergantung sistem operasi 32 atau 64




sekarang kita coba untuk menuliskan angka didalam code
program kita


*/

package main

import "fmt"

func main() {
	fmt.Println("ini angka satu =", 1)
	fmt.Println("ini angka dua =", 2)
	fmt.Println("ini koma =", 3.2)
}

/*

nah sekarang pertanyaannya 1 dan 2 itu type datanya apa?
untuk lebih jelasnya nanti akan kita bahas dimateri variable


*/
