/*

- constant

constant adalah sebuah variable yang nilainya tidak bisa diubah lagi ketika
setelah pertama diberi nilai

cara membuat constant itu mirip dengan variable, akan tetapi menggunakan keyword const
saat pembuatan variable kita wajib secara langsung mengisi datanya apa

kalo kita membuat variable, kita bisa mendeklarasikan variablenya dulu beserta
type datanya, barulah dibaris bawahnya kita bisa mengisikan datanya apa

constant tidak bisa seperti itu, setelah constant dibuat beserta type datanya apa
langsung kita isikan datanya apa

untuk lebih jelasnya langsung saja kita coba




*/

package main

func main() {
	const firstname string = "muiz"
	const lastname = "zuddin"

	// firstname = "hasan"  // akan error

	/*

		jika variable biasa setelah kita buat harus kita pakai
		kalo tidak golang akan mendeteksi adanya error

		kalo constant itu setelah dibuat, tidak apa apa kalo tidak
		kita pakai, karena bisa jadi nanti dipakai difile yang lain

		karena bisanya dipakai dibeberapa file lain, maka
		golang tidak akan menjadikan itu sebagai error


		selain itu constant juga bisa kita deklarasikan sebagai
		multiple constant

	*/

	const (
		umur   int8 = 10
		status bool = true
	)

}
