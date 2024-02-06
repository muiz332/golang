/*

- paraller programming dan concurrency

sebelum kita belajar mengenai goroutine, kita belajar dulu mengenai paraller programming dan
concurrency

saat ini kita hidup diera multicore, dimana jarang sekali kita menggunakan single core processor
dimana sekarang kita bisa dengan mudah membuat proses paraller diaplikasi

paraller progarmming sederhananya adalah memecah suatu masalah dengan cara membaginya menjadi yang lebih kecil
dan dijalankan secara bersamaan dan pada waktu yang bersamaan pula

jadi contohnya kita ada 4 tugas dan prosessor kita memiliki 4 core
maka paraller programming akan menggunakan 4 core tersebut untuk mengerjakan 4 tugas
jadi bisa dikerjakan dalam waktu bersamaan


contoh paraller

menjalankan beberapa applikasi sekaligus disystem operasi kita,
misalkan kita menjalankan browser, visual studio code dan juga word

jadi proses tersebut berjalan secara bersamaan, nah itu disebut dengan paraller


saat kelian belajar mengenai paraller programming kalian akan mengenal 2 hal
yaitu prosess dan theard

prosess
	prosess adalah sebuah eksekusi program
	prosess mengkonsumsi memory yang besar
	prosess saling terisolasi
	prosess lama untuk dijalankan dan dihentikan

thread
	threard adalah segmen atau bagian yang lebih kecil dari prosess yang ada didalam proses itu sendiri
	threard menggunakan memory yang lebih kecil
	threard bisa saling berhubungan jika dalam proses yang sama
	threard cepat dijalankan dan cepat dihentikan


contohnya adakah browser

ketika kalian membuat sebuah browser itu namanya adalah proses
dan didalam browser kalian bisa membuat beberapa tab, tab itu disebut dengan threard

jadi bagian yang lebih kecil yang berada didalam sebuah proses yang sama





- concurrency

lalu apa bedanya dengan concurrency? kalo paraller itu akan menjalankan beberapa tugas secara bersamaan
kalo concurrency itu adalah menjalankan beberapa tugas secara bergantian

dalam paraller kita butuh threard yang banyak sedangkan concurrency kita hanya butuh sedikit threard
dan waktunya tidak bersamaan

dan golang itu sebenarnya bukan paraller jadi defaultnya dia adalah concurrency, tapi karena sekarang kita
menjalankan applikasinya dimulti core, jadi dia sebenarnya campuran antara parller dan concurrency

nanti kita akan bahas lebih detil digoroutine

jadi ketika nanti kalian belajar goroutine, nanti kita akan bahas tentang concurrency
jadi bisa bergonta ganti pekerjaan

tapi tenang semua hal ini tidak akan dikerjakan secara manual dan ini semua sudah dihandle oleh si golang
concurrency



- kapan menggunakan paraller dan concurrency

kalian dapat menggunakan paraller ketika membuat applikasi CPU bound
artinya kalian mengandalkan kecepatan untuk menjalankan applikasi tersebut

contohnya kalo kalian membuat applikasi yang berhubungan dengan mechine learning jadi dia tergantung
dengan kecepatan cpu

karena itu teknologi seperti AI itu banyak sekali yang menggunakan GPU bukan lagi CPU
karena GPU jumlah corenya itu lebih banyak dari pada CPU


lalu kalian dapat menggunakan concurrency ketika membuat applikasi I/O bound
dimana biasanya algoritma atau appilkasi yang sangat ketergantungan dengan kecepatan input dan output
device yang digunakan

rata rata applikasi yang akan kita buat itu kebanyak I/O bound karena bisanya kita akan membaca data
dari file atau juga database

atau mungkin melakukan API call ke service yang lain
jadi seperti web developer, banckend itu rata rata bikin applikasinya seperti ini

kecuali kalian adalah data scienc, mesin learning mungkin kebanyak bikin applikasinya yang CPU bound

nah applikasi yang I/O bound ini walaupun terbantu dengan implementasi paraller programming tetapi benefitnya
akan lebih baik menggunakan concurrency programming

bayangkan kita membaca data dari database dan threard harus menunggu 1 detik karena I/O bound
kalo paraller itu harus nunggu sampai selesai

kalo concurrency itu tidak, kalo dia memanggil database dan ternyata dia membutuhkan waktu 1 detik
dibanding saya menunggu 1 detik, saya akan mengerjakan hal yang lain

ini cuma bisa dilakukan pada bahasa pemrograman yang mendukung concurrency contohnya golang
maka dari itu golang banyak digunakan untuk membuat applikasi backend

karena dia mendukung concurrency programming, dan tidak terlalu menggunakan resource yang terlalu
tinggi karena dia menggunakan concurrency bukan paraller

jadi mudah mudahan kalian paham kenapa kita menggunakan concurrency dan kenapa kita menggunakan
paraller

digolang defaultnya adalah concurrency akan tetapi karena dia juga support multicore artinya
selain concurrency dia juga support tentang paraller programming









*/

package main

func man() {

}
