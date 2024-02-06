/*

- embed file ke string

nah sekarang kita coba untuk embed file ke string
kita coba buat file dulu dengan nama

say.txt

lalu kita baca file tersebut menggunakan package embed
kita coba

*/

package embed_file_string

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed say.txt
var say string

func TestEmbedString(t *testing.T) {

	fmt.Println(say)

	/*

		jadi kalo saya jalankan maka hasilnya seperti ini

		hello selamat datang
		apa kabar

		hasilnya sesuai ya dengan text yang ada didalam say.txt
		jadi embed ini akan sering kita gunakan saat membuat applikasi web berbasis golang

		jadi seperti itu ya cara embed file ke


	*/
}
