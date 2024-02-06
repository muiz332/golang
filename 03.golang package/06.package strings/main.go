/*


- package strings

kali ini kita akan coba package strings, jadi package strings ini digunakan untuk memanipulasi
string

jadi sebenarnya ada banyak function untuk memanipulasi string, akan tetapi
kita akan bahas beberapa saja ya

oke langsung saja kita coba

*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	// toLower()

	fmt.Println(strings.ToLower("Hello World"))                    // mengubah menjadi huruf kecil
	fmt.Println(strings.ToUpper("Hello World"))                    // mengubah menjadi huruf besar
	fmt.Println(strings.Trim("    Hello World  ", " "))            // mengilangkan karakter awal dan akhir
	fmt.Println(strings.ReplaceAll("Hello World", "Hello", "hai")) // mengubah string

	/*

		jadi seperti itu saja ya penggunaan dari package string
		sangat simple dan sederhana

		mudah mudahan kalian paham

	*/
}
