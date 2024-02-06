/*

- buffered channel


seperti yang sudah dijelaskan sebelumnya bahwa channel itu hanya bisa menerima 1 data saja
artinya jika kita menambahkan data ke 2 maka kita akan diminta menunggu sampai data 1 itu
ada yang mengambil

kadang ada kasus dimana pengirim lebih cepat dibandingkan penerima, dalam hal ini kita menggunakan
channel maka secara otomatis pengirim akan ikut lambat

karena jika penerimanya lambat memproses datanya maka pengirim yang terlalu cepat dia
akan menunggu sampai datanya diambil sama penerima

akhirnya ketika penerima lambat pengirim ikut lambat

jadi buffered channel itu adalah cara untuk menampung lebih banyak data didalam channel

kalian juga perlu tahu kalo didalam buffer itu ada istilahnya capacity
jadi buffer itu kyk penyimpanan yang ada didalam channel

dan kita bebas memasukkan berapa jumlah kapasitas antrian didalam buffer
kita kita set 5, maka kita hanya bisa menerima 5 data dibuffer

dan jika kita menerima data ke 6 maka data yang 5 tadi harus ada yang mengambil terlebih dahulu
ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari pada pengirim

oke langsung saja kita coba






*/

package main

func main() {

	var channel chan string = make(chan string, 3)
	defer close(channel)

	channel <- "selamat datang"
	channel <- "halo"
	channel <- "selamat pagi"
	channel <- "selamat pagi"

	/*

		kalo kita jalankan maka tidak akan error
		padahal belum ada yang menerima, karena ketika kita mengirim 3 buah data itu tidak akan error
		karena masuknya kedalam buffer

		kalo kita tidak menggunakan buffer maka akan terjadi error karena tidak ada yang menerima
		akan tetapi kalo kita kirim 1 kali lagi data yang ke 4

		maka akan terjadi error karena ketika kita mengirim data ke 4, maka data yang sebelumnya
		harus ada yang menerima dulu

		artinya channelnya harus kosong dulu baru boleh mengirim data yang ke 4
		jadi gitu ya buffered channel

		mudah muahan kalian paham

	*/

}
