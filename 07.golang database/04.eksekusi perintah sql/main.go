/*


- eksekusi perintah sql


nah kita akan bahas dulu eksekusi sql yang bukan query, jadi seperti insert, update, delete, create table
dan apapun yang tidak membutuhkan pengambalian data ya

golang telah menyediakan sebuah function yang namanya, db.ExecContext(context,sql,params)

ketika kita mengirim perintah sql kita butuh mengirimkan context dan dengan context
kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengirimkan perintah sqlnya


*/

package main

import (
	"context"
	"fmt"
	"sql-nonQuery/db"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	mysqlConn := db.CreateConnection()
	defer mysqlConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := mysqlConn.ExecContext(ctx, "INSERT INTO orang (nama, alamat) VALUES ('hasan','banyuwangi');")
	if err != nil {
		panic(err)
	}
	fmt.Println("berhasil menambahkan data ke database")

	/*

		jadi struktur table yang saya buat seperti ini
		+--------+-------------+------+-----+---------+----------------+
		| Field  | Type        | Null | Key | Default | Extra          |
		+--------+-------------+------+-----+---------+----------------+
		| id     | int(11)     | NO   | PRI | NULL    | auto_increment |
		| nama   | varchar(20) | NO   |     | NULL    |                |
		| alamat | varchar(50) | NO   |     | NULL    |                |
		+--------+-------------+------+-----+---------+----------------+

		kalo kalian jalankan dan kalian lihat kita sudah berhasil memasukkan datanya
		kedalam database


		+-----+------+------------+
		| id  | nama | alamat     |
		+-----+------+------------+
		| 180 | muiz | banyuwangi |
		+-----+------+------------+
		1 row in set (0.011 sec)


		pada contoh diatas kita juga menggunakan contextWithTimeout atau kalian hanya menggunakan
		context.background saja juga bisa

		oke jadi seperti itu ya bagaimana kita berinteraksi dengan mysql
		selanjutnya kita akan coba melakukan query

	*/
}
