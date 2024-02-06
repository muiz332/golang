/*

- membuat koneksi ke database


jadi hal pertama sebelum kita membuat applikasi nya kita harus melakukan koneksi dulu
kedatabasenya

agar nanti kita bisa berinteraksi dengan database yang ingin kita gunakan
untuk melakukan koneksi kedatabase kita bisa membuat object sql.DB dengan menggunakan
function sql.Open(driver, datasourcename)

untuk menggunakan mysql, kita bisa menggunakan driver "mysql"
sedangkan data source name tiap database memiliki cara sendiri sendiri

untuk mysql kalian bisa tulis seperti ini

user:password@tcp(localhost:5555)/dbname

jadi username password database kita apa, lalu kita bisa gunakan localhost beserta port database
untuk berinteraksi dengan database dikomputer kita

dan juga kita harus buat dulu database namenya dimysql nya nanti
kemudian kita tuliskan di /dbname


jika object sql.Db tidak digunakan lagi maka disarankan untuk diclose agar tidak menyebabkan
connection leak

sebelum itu kalian bikin nama database misal test atau sesuai yang kalian inginkan
oke langsung saja kita coba




*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("database berhasil terkoneksi")

	/*

		kalo kalian jalankan dan tidak terjadi error, maka kalian sudah berhasil
		terkoneksi dengan database mysqlnya

		dan jangan lupa kalian gunakan defer untuk melakukan close connectionnya
		agar nanti setelah applikasi kita selesai dijalankan

		maka semua connectionnya akan diclose

		dan kalian jangan lupa untuk mengimport 	_ "github.com/go-sql-driver/mysql"
		karena kita butuh untuk menjalankan function init yang ada didalam package tersebut

		oke jadi seperti itu ya bagaimana cara kita terkoneksi dengan mysql
		mudah mudahan kalian paham

	*/
}
