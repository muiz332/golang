/*

- type data map

kali ini kita akan mambahas type data map digolang


apa itu type data map?
type data map adalah sebuah type data yang menampung kumpulan data
dimana data tersebut dapat kita akses dengan type data selain integer\

untuk jumlah data yang bisa ditampung itu tidak terbatas ya
tidak seperti array yang ada batasnya, kalian bisa menambahkan
berapapun didalam map ini

kalo array kemarin kan untuk mengakses datanya kita bisa menggunakan
index, nah kalo map kita bisa menggunakan type data yang lain
asalkan tidak boleh sama

atau simplenya pasangan antara key dan value, akan tetapi
kalo ada key yang sama, maka key tersebut akan ditimpa dengan
key dan value yang baru

langsung saja ya kita coba

*/

package main

import "fmt"

func main() {

	var person map[string]string = map[string]string{
		"nama":    "muiz",
		"jurusan": "teknik informatika",
	}

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["jurusan"])

	/*

		kalo kalian ingin menambahkannya diluar kurung kalian bisa
		tulis seperti ini

	*/

	person["alamat"] = "banyuwangi"
	fmt.Println(person["alamat"]) // banyuwangi

	/*

		atau kalian juga bisa mengubah valuenya,
		sama seperti array, akan tetapi yang kalian butuhkan adalah
		keynya kemudian kita assigment dengan value yang baru

	*/

	person["nama"] = "muizzuddin"
	fmt.Println(person["nama"])

	/*

		selanjutnya ada beberapa function yang dapat kita gunakan
		untuk berinteraksi dengan type data map

		len()
		make(map[typekey]typevalue)
		delete(map,key)

		kita coba

	*/

	fmt.Println(len(person)) // 3

	// make

	orang := make(map[string]string)
	orang["nama"] = "hasan"
	orang["jurusan"] = "agama"
	fmt.Println(orang)

	// delete
	delete(orang, "jurusan")
	fmt.Println(orang)
}
