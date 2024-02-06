/*

- embed hasil dicompile

hasil embed yang dilakukan oleh package embed ketika dicompile itu akan permanent
jadi beneran file external ketika diembed akan menjadi byte dan byte tersebut
akan menjadi bagian dari code program kita

jadi setelah kalian compile, file external yang kalian embed itu kalian hapus
itu tidak akan menjadi masalah


karena kan bytenya sudah ada didalam code program kita
jadi byte tersebut akan secara permanent masuk kedalam variablenya

oke sekarang kita coba



*/

package main

import (
	_ "embed"
	"fmt"
)

//go:embed say.txt
var say string

func main() {

	fmt.Println(say)
}
