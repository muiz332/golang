/*



- pointer function new

sebelumnya kita membuat pointer dengan tanda & diikuti dengan nama variablenya
atau data yang baru

dengan menggunakan function new kita bisa membuat  pointer yang isinya data kosong
artinya tidak ada data awal

kita coba

*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var alamat1 *Address = new(Address)
	alamat1.City = "banyuwangi"
	fmt.Println(alamat1) // &{banyuwangi  }

	/*

		jadi kita bisa membuat pointer kosong dengan menggunakan
		function new

		jadi seperti itu pointer digolang

	*/

}
