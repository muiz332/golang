/*

- function parameter

dimateri kali ini kita akan bahas mengenai parameter

saat kita membuat function kadang kadang kita membutuhkan data dari luar
atau kita sebut dengan parameter

kita bisa menambahkan parameter lebih dari satu,
dan kalo kita menambahkan parameter didalam function, maka kita wajib
menambahkan argument saat function tersebut dipanggil

argument itu kita memasukkan data didalam kurung buka dan tutup
saat function dipanggil

langsung saja ya kita coba


*/

package main

import "fmt"

func sayHello(nama string) {
	fmt.Println("hello nama saya", nama)
}

func main() {
	sayHello("muiz")  // hello muiz
	sayHello("hasan") // hello hasan

	/*

		jadi kita bisa menambahkan parameter atau simplenya variable didalam
		kurung kurawal buka dan tutup setelah nama functionnya beserta type datanya

		setelah itu saat function dipanggil, maka kita harus mengirimkan data tersebut
		dan data tersebut akan dimasukkan kedalam variable nama, variable nama inilah
		yang disebut dengan parameter

		jadi kita bisa membutuhkan data dari luar function yang dapat kita kelola
		didalam function

		jadi code kita bisa lebih dynamis atau bisa menyesuaikan data apa yang
		dikirimkan kedalam parameternya

		dan kalo lebih dari satu parameter kalian tinggal pisahkan saja
		dengan koma

		jadi ini lah function parameter




	*/
}
