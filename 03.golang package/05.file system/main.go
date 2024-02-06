/*


- file system

dimateri kali ini kita akan belajar bagaimana berinteraksi dengan file system
digolang

apa saja yang akan kita lakukan

- membuat folder
- menghapus folder
- membuat file
- menghapus file
- membaca file

demgan menggunakan package os kita bisa melakukan hal itu
oke langsung saja kita coba

*/

package main

import (
	"fmt"
	"os"
)

func main() {

	// membuat folder

	var errMkdir error = os.Mkdir("folder baru", 0750)
	if errMkdir != nil {
		fmt.Println(errMkdir.Error())
	}

	// os.Mkdir("folder baru", 0750)

	/*

		kita bisa menggunakan function Mkdir dengan 2 buah parameter
		parameter  pertama itu nama foldernya dan parameter kedua itu adalah hak aksessnya
		untuk hak akses tersebut berlaku pada sistem operasi linux sedangkan windows berbeda

		tetapi kita tetap menuliskan parameternya, kalian tuliskan saja 0750
		seperti itu ya

		0: Tidak ada izin (tidak dapat membaca, menulis, atau menjalankan).
		1: Hanya izin eksekusi (biasanya digunakan untuk file eksekusi).
		2: Hanya izin tulis.
		3: Izin tulis dan eksekusi.
		4: Hanya izin baca.
		5: Izin baca dan eksekusi.
		6: Izin baca dan tulis.
		7: Izin baca, tulis, dan eksekusi (maksimum izin).

		kalo kita jalankan maka akan membuat sebuah folder dengan nama folder baru
		akan tetapi jika foldernya sudah ada, kalo kita jalankan lagi
		maka itu akan terjadi error
	*/

	// menghapus folder

	os.Mkdir("baru", 0750)
	os.Remove("baru")

	/*

		kita cukup menggunakan function remove setelah itu diikuti dengan path dari filenya ya
		jadi sederhana banget ya

		nah selanjutnya kita akan coba untuk membuat file
		misalkan saya ingin membuat file yang berada didalam folder baru

		kita coba

	*/

	os.WriteFile("./folder baru/test.txt", []byte("ini adalah text \nbaris selanjutnya"), 0666)

	/*

		kita gunakan function WriteFile dimana terdapat 3 buah parameter
		parameter pertama itu adalah lokasi file yang akan kita buat yaitu didalam folder baru
		parameter kedua itu adalah data apa yang mau kita tuliskan kedalam file yang kita buat

		dimana datanya dalam bentuk byte, jadi byte itu alias dari type ada uint8
		dengan menggunakan uint8 ini maka dapat mencakup semua karakter

		karena kita perlu mengubah semua karakter menjadi unicode terlebih dahulu
		dimana kita tampung kedalam sebuah slice

		jadi sebenarnya sama seperti ini
		byte('a') menggunakan kutip satu untuk mengubah satu karakter menjadi unicode

		kalo ingin mengubah lebih dari satu karakter kita bisa gunakan slice yang didalamnya terdapat
		type data byte, jadi kita bisa gunakan kutip 2

		dan terakhir adalah parameter ke 3 yaitu untuk mengaktur hak akses, tetapi kalian jangan
		terlalu memikirkan hal itu terlebih dahulu, kalian bisa menuliskan 0666 saja ya

		sekarang kita coba remvove file

	*/

	os.WriteFile("file-dihapus.txt", []byte("muiz"), 0666)
	os.Remove("./file-dihapus.txt")

	/*

		jadi kita bisa menggunakan function Remove, jadi function remove ini juga bisa kita gunakan
		untuk menghapus sebuah folder ataupun file

		sekarang kita akan coba membaca file

	*/

	var text, errReadFile = os.ReadFile("./folder baru/test.txt")
	if errReadFile == nil {
		fmt.Println(string(text))
	} else {
		fmt.Println(errReadFile.Error())
	}

	/*

		jadi kita menggunakan function ReadFile dengan menambahkan parameter letak dari file tersebut
		function tersebut mengembalikan []byte dan error

		kalo file tidak ada maka terjadi error, dan kalo ada maka kita akan mendapatkan data
		byte dan kita harus mengconversi menjadi string, caranya kalian cukup tuliskan
		type data string kurung buka dan tutup dan masukkan parameter textnya

		meskipun slice yang kita masukkan kedalam parameternya, tetap bisa diconversi oleh
		string() nya

		jadi gitu ya mengenai file system
		mudah muahan kalian paham

	*/

}
