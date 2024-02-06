/*

- context with value


pada saat awal kita membuat context, context itu tidak memiliki value yang kita buat pakai
background atau todo

nah kita bisa manambahkan data yang berupa key value ke dalam context
saat kita menambah value ke context secara otomatis akan membuat child context baru
artinya context aslinya itu tidak akan berubah

untuk menambahkan value ke context kita bisa menggunakan function
context.WithValue(parent, key, value)

seperti yang sudah kita bahas ya, kalo misalkan kita menambah, atau mengubah sesuatu pada
context, maka secara otomatis akan membuat child context baru

oke kita coba bikin context seperti yang sudah dijelaskan kemarin  ya
ada context A beserta childnya

kita coba

*/

package main

import (
	"context"
	"fmt"
)

func main() {

	var contextA = context.Background()

	var contextB = context.WithValue(contextA, "b", "B")
	var contextC = context.WithValue(contextA, "c", "C")

	var contextD = context.WithValue(contextB, "d", "D")
	var contextE = context.WithValue(contextB, "e", "E")

	var contextF = context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	/*

		kalo kalian jalankan maka munculnya seperti ini

		context.Background
		context.Background.WithValue(type string, val B)
		context.Background.WithValue(type string, val C)
		context.Background.WithValue(type string, val B).WithValue(type string, val D)
		context.Background.WithValue(type string, val B).WithValue(type string, val E)
		context.Background.WithValue(type string, val C).WithValue(type string, val F)

		jadi kalo misalkan kita lihat context F, hasilnya itu dari context A memiliki child
		ke context C dan context C memiliki child context F

		jadi ini bacanya saya akan membuat context B yang memiliki parent ke context A
		dimana context B memiliki key dan value b, B

		jadi seperti itu ya cara menambahkan data ke dalam context
		dan key valuenya itu type datanya interface kosong ya, artinya kalian bisa
		bebas memasukkan type data apapun


		nah sekarang kita coba untuk mendapatkan value dari context yang sudah kita buat
		untuk mengambil datanya kalian bisa gunakan method value dengan parameter keynya

		tetapi kalian harus ingat bahwa ketika data yang kita ambil dari context itu
		tidak ada, maka golang akan mencari ke parent contextnya

		ketika parent contextnya tidak ada, maka dia akan mencari ke parentnya lagi
		terus seperti itu sampai ketemu, kalo tidak ketemu dia akan mengembalikan nil

		jadi golang mencarinya ke atas ya bukan ke bawah childnya
		oke langsung saja kita coba

	*/

	fmt.Println()
	fmt.Println(contextF.Value("f")) // F
	fmt.Println(contextF.Value("c")) // C
	fmt.Println(contextF.Value("b")) // nil
	fmt.Println(contextA.Value("b")) // nil

	/*

		kalo kita jalankan maka akan ada dua buah nilai nil
		karena seperti yang sudah dijalaskan bahwa golang mencari value didalam context
		itu dari bawah keatas ya

		jadi kalo kita lihat contextF itu mencari key b, dimana key b itu bukan
		parent dari contextF melainkan saudara dari parentF

		maka nilainya adalah nil
		sedangkan contextA itu mencari key b, dimana key b itu ada didalam child contextB

		jadi hasilnya akan nil juga karena golang mencari key dari child menuju parent
		bukan dari parent kechild

		seperti yang kemarin dijalaskan bahwa parent itu membawa valuenya ke childnya
		tetapi sebenarnya bukan membawa tetapi lebih ke arah hubungan ya, saling terhubung

		oke jadi seperti itu ya context with value
		mudah mudahan kalian paham



	*/

}
