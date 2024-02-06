/*


- membuat context

karena context itu adalah interface, untuk membuat context kita harus membuat struct dulu yang sesuai
dengan interface contextnya

akan tetapi golang sudah menyediakannya jadi kita tidak perlu membuat structnya secara manual
digolang package context terdapat function yang dapat kita gunakan untuk membuat context

ada dua function yang kita gunakan untuk membuat context



- context.Background()

membuat context kosong, tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki value apapun
biasanya digunakan di main function atau didalam test, atau diawal proses request terjadi

- context.TODO()

membuat context kosong seperti context.Background(), namun biasanya digunakan ketika belum jelas context
apa yang ingin digunakan

oke langsung saja kita coba



*/

package main

import (
	"context"
	"fmt"
)

func main() {

	// membuat context background

	var background context.Context = context.Background()
	fmt.Println(background)

	// membuat context todo

	var todo context.Context = context.TODO()
	fmt.Println(todo)

	/*

		kalo kita jalankan maka hasilnya seperti ini

		context.Background
		context.TODO

		jadi hanya muncul string ya
		tetapi kalo nanti kalian sudah bikin applikasi website menggunakan golang web
		itu kalian tidak perlu bikin context secara manual

		nanti otomatis didalam setiap request yang diterima oleh applikasi kita
		secara otomatis nanti akan dibuatkan contextnya

		jadi kita tinggal pakai saja
		jadi ini tentang context

		mudah mudahan kalian paham

	*/

}
