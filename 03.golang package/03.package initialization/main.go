/*


- package initialization


saat kita membuat package kita bisa membuat sebuah function yang akan dijalankan ketika
package itu dipanggil

ini sangat cocok ketika kita ingin mengkoneksikan dengan database
untuk membuat function yang bisa secara otomatis dijalankan saat package dipanggil

kita cukup membuat function dengan nama ini

kita coba

*/

// package main

// import "package-initialization/db"

// func main() {
// 	db.GetConn()

// }

/*

	kalo kita jalankan maka akan tampil mysql
	padahal kita tidak menjalankan function init

	ini berarti ketika kita memanggil package, maka function init akan dijalankan terlebh dahulu
	lalu bagaimana kalo kita memiliki 2 buah file yang sama sama memiliki function init ?

	khusus untuk function dengan nama init, kita bisa buat lebih dari satu function init
	pada file yang berbeda

	jadi memeskipun ketik tidak memanggil sebuah function yang ada didalam file
	mysqlConn.go

	secara otomatis function init tersebut akan dijalankan dengan urutan file yang paling teratas
	akan tetapi kalo kalian hanya melakukan import pada package dan tidak menggunakan
	package tersebut maka golang akan mendeteksi error

	function init dijalankan ketika kita memanggil package
	golang akan menganggap error ketika package yang dipanggil tidak digunakan

	adakah cara bagaimana golang bisa mengabaikan package yang kita panggil meskipun
	tidak kita gunakan

	kita bisa menggunaka blank identifier
	jadi kalian cukup menambahkan tanda underscore sebelum nama packagenya


*/

package main

import _ "package-initialization/db"

func main() {
}

/*

jadi kalian tinggal nemambahkan underscore maka package yang tidak digunakan tidak
akan dianggap error oleh golang

jadi ceritanya kalian ingin mengkoneksikan dengan database
kalian bisa menggunakan fitur package initialization ini

jadi seperti itu ya mengenai package initialization
mudah mudahan kalian paham

*/
