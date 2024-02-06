/*

- golang context


dimateri kali ini kita akan belajar mengenai golang context
apa itu context?

context adalah sebuah data yang membawa value, sinyal cancel, sinya timeout,
dan sinya deadline

context biasanya dibuat per request, misalnya setiap ada request yang masuk ke server web
melalui http request

kalo kalian nanti belajar golang web, nanti otomatis akan ada contextnya per request yang masuk
kalo sekarang kita akan bikin contextnya secara manual dulu ya

context digunakan untuk mempermudah kita meneruskan value dan sinya antar proses

kalo kita menggunakan golang nanti akan banyak sekali menggunakan goroutine, kadang kita butuh
kalo ada proses dibatalkan maka semua proses ikutan dibatalkan

nah dengan menggunakan context hal ini itu bisa dilakukan, kalo sebelumnya kan kalo goroutinenya sudah
berjalan sendiri kita tidak bisa apa apa

nah dengan adanya context itu kita bisa mengirimkan sinya cancel untuk membatalkan contextnya
dan goroutinenya nanti bisa tahu ternyata ada sinya cancel dari context

maka secara otomatis goroutinenya akan membatalkan prosesnya, dan itu akan kita pelajari seiring materi
jadi ini adalah context



- kenapa context perlu dipelajar ?

context digolang bisa digunakan untuk mengirim data request atau sinyal ke proses lain atau dari
goroutine ke goroutine yang lain

dengan menggunakan context, ketika kita ingin membatalkan semua proses, kita cukup mengirim
sinyal ke context, maka secara otomatis semua proses dibatalkan

hampir semua bagian digolang memanfaatkan context, seperti database, http server, http client,
dan lain lain

bahkan digoogle sendiri, ketika menggunakan golang, context wajib digunakan dan selalu dikirim ke setiap
function yang dikirim

jadi hampir semua function function digoogle sendiri, mereka saat bikin function function pasti ada
parameter contextnya

itu untuk memastikan bahwa function function tersebut bisa dibatalkan dan lain lain
jadi context ini sangat wajib untuk dipelajari

dan akan mempermudah kita untuk membuat applikasi yang lumayan komplex



- cara kerja context

					-------------
                    | proses A	|
					|			|
					| context	|
					-------------
										dikirim ke

										-------------
										| 			|
										| Proses B	|
										|        	|
										-------------

dikirim ke

-------------
| 			|
| Proses C	|
|        	|
-------------




misalkan kita punya proses A yang memiliki context
saat kita membuat proses yang lain contohnya proses B dan proses C atau goroutine

nah context ini bisa dikirim ke goroutine atau proses yang lain
nanti ketika kita bilang ke proses A untuk membatalkan contextnya

maka sinyalnya akan dikirim ke semua proses yang memiliki context dari proses A akan
ikut dibatalkan



- package context

jadi context itu direpresentasikan kedalam sebuah interface Context
interface context terdapat didalam package context

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

jadi sebenarnya interface context isinya hanya 4 method saja
yang nanti akan kita bahas satu persatu ya

oke mudah mudahan kalian bisa mengikuti materi golang
context ini sampai akhir ya

dimateri selanjutnya kita akan bahas bagaimana cara membuat context



*/

package main

func main() {

}
