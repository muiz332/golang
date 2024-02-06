/*


- switch

switch itu sama seperti if, akan tetapi kalo if kan kita bisa masukkan
kondisi sekomplex mungkin

kita bisa memasukkan operasi perbandian, lebih kecil dari lebih besar dari
dan lain lain

nah switch hanya bisa menggunakan operasi perbandingan sama dengan saja
jadi kalo kalian ingin membuat kondisi

dan kondisinya itu hanya menggunakan operasi sama dengan, maka gunakanlah
switch

langsung saja kita coba ya


*/

package main

import "fmt"

func main() {

	var nama string = "hasan"

	switch nama {
	case "muiz":
		fmt.Println("hello muiz")
		fmt.Println("hello muiz")
		fmt.Println("hello muiz")
	case "husain":
		fmt.Println("apa kabar", nama)
	default:
		fmt.Println("siapa anda")
	}

	/*

		jadi seperti itu sintaxnya ya, kalian tinggal menggunakan keyword switch
		dan kemudian variable yang mau dicheck

		case itu anggap saja operasi sama dengan ya
		jadi kalo case muiz

		maka yang dijalankan akan print hello muiz

		untuk default itu sama dengan else, jadi ketika kondisinya semua bernilai
		false, dan kita ingin menangkap yang kondisinya false

		maka gunakan default, kalo kalian tidak gunakan juga tidak apa apa

		dan selanjutnya ada switch dengan short statement
		jadi kita bisa deklarasikan variable nya setelah switch nya


	*/

	switch jurusan := "teknik informatika"; jurusan {
	case "teknik informatika":
		fmt.Println("hello")
	default:
		fmt.Println("apa kabar")

	}

	/*

		jadi kita bisa membuat variable setelah switchnya ya
	*/

}
