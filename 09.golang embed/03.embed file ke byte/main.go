/*

- embed file ke []byte

ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary
seperti gambar, video, audio atau file file lain yang tidak bisa
dibaca oleh editor kita

oke langsung saja kita coba



*/

package main

import (
	_ "embed"
	"io/fs"
	"os"
)

//go:embed gambar.jpg
var gambar []byte

func main() {
	err := os.WriteFile("gambar_baru.jpg", gambar, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	/*

		jadi saya baca filenya dan saya akan buat file baru dengan data dari file yang
		saya baca dengan embed tadi

		jadi kalo saya jalankan maka akan menghasilkan file gambar yang sama
		dengan file gambar yang saya baca dengan package embed

		jadi seperti itu ya
		mudah mudahan kalian paham

	*/
}
