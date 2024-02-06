/*


- function

dimateri kali ini kita akan bahas mengenai function

sebenarnya kalian sudah sering melihat function
yatiu ketika kalian menuliskan

func main(){

}

itu adalah function, akan tetapi kita juga bisa membuat function sendiri


jadi apa itu function ?
function adalah sebuah code program yang dapat kita panggil berulang ulang kali
dan code program tersebut menjalankan satu tugas

untuk membuat function kalian bisa tuliskan seperti ini

func namaFunctionnya (){

}

jadi nama functionnya menggunakan aturan penulisan camesCase
yaitu huruf pertama dikata awal huruf kecil, dan huruf pertama dikata kedua
dan seterusnya itu menggunakan huruf besar

kalo kalian hanya membuat functionnya saja, tanpa dipanggil
code yang ada didalam function tidak akan dijalankan

maka kita harus panggil functionnya dulu, caranya
kalian tuliskan nama functionnya, diikuti dengan kurung buka dan tutup

kita coba


*/

package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() {

	// sekarang kita panggil functionnya
	sayHello()

	/*

		kalo kita jalankan maka akan muncul hello
		nah misalkan saya ingin memanggil functionnya lagi

		jadi saya hanya menuliskan nama functionnya
		dan kurung buka tutup

	*/

	sayHello()
	fmt.Println()

	/*

		kalo saya jalankan
		maka akan ada 2 hello yang tampil kelayar

		bayangkan kalo kalian memiliki sebuah code program yang
		dapat mengerjakan satu tugas tertentu

		dan code program tersebut digunakan berkali kali
		kita bisa buat code program tersebut menjadi sebuah function

		jadi nanti baris codenya tidak akan panjang lagi
		karena mungkin function tersebut kita bisa buat difile lain

		dan kita panggil difile yang kita butuhkan, selain itu penggunaan function
		itu agar code kita lebih modular

		dalam artian, code kita dapat dipisah pisahkan sesuai dengan tugasnya
		masing masing

		agar lebih mudah dibaca, karena semakin baik code program kalian
		maka akan mudah dibaca

		jadi code tersebut bisa dipanggil berapapun ya
		atau misalkan kita coba menggunakan perulangan
		untuk memanggilnya

	*/

	for i := 0; i < 5; i++ {
		sayHello()
	}

	/*

		maka function tersebut akan dijalankan sebanyak 5 kali
		jadi seperti itu ya penggunaan function

		selanjutnya akan akan membahas function lebih detil lagi

		didalam golang, karena dia bukan bahasa pemrograman yang Object Oriented Programing
		maka kita akan sering menggunakan function digolang

		jadi function ini sangat penting, dan harus kalian pahami
		karena kedepannya kita akan sering menggunakan function



	*/
}
