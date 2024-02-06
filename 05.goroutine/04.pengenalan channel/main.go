/*

- channel


dimateri kali ini kita akan belajar mengenai channel

channel adalah tempat komunikasi secara syncronous yang bisa dilakukan oleh goroutine
jadi kita tahu kalo gorutine itu kita coba memberikan return value
maka return value tersebut tidak akan bisa ditangkap karena dia masih pending

nah sekarang pertanyaaanya gimana caranya kita bisa menangkap return value tersebut
tentu saja kita harus menunggu pendingnya tersebut menggunakan channel

jadi channel ini adalah komunsikasi secara syncronous antara satu atau lebih goroutine
jadi kalo kalian ingin mengirim data dari goroutine ke goroutine yang lain
kalian bisa menggunakan channel

didalam channel itu terdapat yang namanya pengirim dan penerima, biasanya pengirim dan penerima
adalah goroutine yang berbeda

jadi jarang banget digunakan pada gorutine yang sama meskipun itu bisa saja
saat kita melakukan pengiriman data ke channel

maka gorutine akan terblock artinya goroutinenya tidak akan jalan sampai sampai data yang
ada didalam channel itu mengambilnya

maka dari itu channel disebut dengan alat komunikasi yang syncronous
nah channel itu cocok sekali sebagai alternatif dari async dan await


jadi kalo kalian sudah belajar mengenai javascript maka goroutine dan channel itu mirip
dengan async dan await pada javascript



- karakteristik dari channel

secara default channel itu bisa menampung satu data, jika kita ingin menambahkan data lagi
harus menunggu data yang ada dichannel itu diambil tersebih dahulu
barulah kita bisa mengambil data selanjutnya dari channel
jadi ini secara bergantian

channel hanya bisa menerima satu jenis data
jadi kalo kalian sudah menentukan type datanya maka type data yang ada didalam channel
itu sesuai dengan apa yang kita tuliskan

channel bisa diambil lebih dari satu gorotuine
channel harus diclose jika tidak digunakan, atau bisa menyebabkan memory leak (data yang tersimpan
dimemory tetapi tidak digunakan yang akan terjadi penumpukan)

dan kalo sudah diclose, pengirim atau penerima akan ditolak

dimateri selanjutnya kita akan bahas bagaimana cara membuat channel



*/

package main

func main() {

}
