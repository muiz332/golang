/*

- path matcher

kemarin kalo kita menggunakan multiple embed itukan harus manual satu satu
menuliskan go:embednya

nah dengan menggunakan path matcher kita bisa memilih satu atau semua file
yang kita ingin embed


*/

package main

import (
	"embed"
	"fmt"
)

//go:embed files/*.txt
var files embed.FS

func main() {

	dir, _ := files.ReadDir("files") // mengambil path didalam directory files
	for _, entry := range dir {

		if !entry.IsDir() { // mengecheck apakah pathnya itu file atau directory

			content, _ := files.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}

	/*

		jadi seperti itu ya, kita bisa gunakan looping untuk membaca semua file

		dengan menggunakan *.txt kita bisa membaca seluruh file dengan extention .txt
		jadi seperti itu tentang path matcher

		mudah mudahan kalian paham

	*/

}
