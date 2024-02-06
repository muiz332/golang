/*
- pengenalan goroutine


goroutine itu adalah sebuah thread ringan yang dikelola oleh goruntime
sebenarnya bukan thread ringan juga, karena alanoginya agak susah biasanya
orang orang menyebutkan thread ringan

jadi goroutine itu sebenarnya dia berjalan didalam thread
ukuran goroutine itu sangat kecil sekitar 2kb jauh lebih kecil dibandingkan dengan thread

thread bisa sampai 1mb atau 1000kb, namun tidak seperti thread yang berjalan secara
paraller akan tetapi goroutine berjalan secara concurrency

jadi ukuran goroutine itu kecil sekali, makanya disebut dengan thread ringan
karena kita bisa buat 500 gorutine yang setara dengan 1 thread

makanya golang itu terkenal dengan minim resource karena untuk membuat goroutine saja
hanya butuh 2kb

jadi kan goroutine itu berjalan didalam thread, nah thread itu lah yang berjalan secara paraller
sedangkan goroutine yang ada didalam thread itu berjalan secara concurrency

jadi karena itulah digolang disebutnya concurrency

jadi gorotinenya itu kalo berjalan akan seperti ini

func proses panjang
func proses1
func proses2

maka ketika bertemu dengan prosess panjang akan dijalankan terlebih dahulu ketika lama akan distop dan dilanjutkan keproses selanjutnya
lalu balik lagi ke proses yang panjang, kalo dijalankan masih lama maka akan distop lagi dan akan lanjut keproses selanjutnya
kemudian balik lagi ke goroutinenya lalu prosesnya dijalankan dan ternyata tidak membutuhkan waktu yang lama
program selesai dijalankan

maka akan jadi seperti ini

func proses1
func proses2
func proses panjang

nah ini mirip dengan konsep asyncronous ya, jadi dia tidak akan menunggu sampai proses yang
panjang tersebut selesai, dia akan melanjutkan ke prosess yang selanjutnya

sebenarnya goroutine itu dijalankan oleh go scheduler dialam thread dimana jumlah threadnya itu sebanyak
GOMAXPROCS (biasanya sejumlah core cpu)

nah dengan itu kita tidak perlu melakukan management thread secara manual karena
semua sudah diatur didalam go scheduler didalam go runtimenya

nah didalam go scheduler itu ada 3 istilah

G  : Goroutine
M  : Thread (Mechine)
P  : Processor


jadi tenang saja kita tidak perlu mengatur bagaimana perpindahan goroutinenya
semua sudah diatur oleh go scheduler didalam goruntime

mudah muahan kalian paham



*/

package main

func main() {

}
