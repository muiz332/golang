/*

- anonymous function


apa itu anonymous function
adalah sebuah function yang tidak memiliki nama

ketika kita menggunakan function sebagai value, ada kalanya itu lebih mudah
jika kita langsung memasukkan functionnya kedalam variable atau parameter
tanpa nama

hal itu disebut dengan anonymous function

langsung saja ya kita coba

*/

package main

import "fmt"

func cetakAndCheck(nama string, check func(string) string) {
	var result string = check(nama)
	fmt.Println(result)
}

func main() {

	cetakAndCheck("muiz", func(str string) string {
		if len(str) == 0 {
			return "nama tidak boleh kosong"
		}
		return str
	})

	cetakAndCheck("", func(str string) string {
		if len(str) == 0 {
			return "nama tidak boleh kosong"
		}
		return str
	})

	/*

		jadi dengan begini sebenarnya logic jadi function yang dijadikan
		sebagai parameter itu bisa dinamis, bisa berubah ubah sesuai
		dengan kebutuhan

		jadi itulah anonymous function
		mudah mudahan kalian paham

	*/

}
