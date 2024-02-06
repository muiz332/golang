/*

- embed multiple file

nah sekarang kita akan coba melakukan embed multiple file ya
kdang kita ingin melakukan embed beberapa file sekaligus

hal ini kita bisa lakukan dengan menambahkan //go:embed dan
menggunakan type data embed.FS

oke langsung saja kita coba


*/

package main

import (
	"embed"
	"fmt"
)

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt

var files embed.FS

func main() {

	// membaca multiple file yang kita embed

	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))

	/*

		kalo kita jalankan maka akan seperti ini ya

		ini file a
		ini file b
		ini file c

		jadi seperti itu cara kita membaca multiple file menggunakan package
		embed

		mudah mudahan kalian paham

	*/

}
