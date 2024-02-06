/*


- sql injection


saat membuat applikasi, kita tidak mungkin melakukan harcode atau menulis secara langsung
perintah sql pada golang

biasanya kita akan menerima input data dari user, lalu membuat perintah sql dari input user
dan mengirimnya menggunakan perintah sql

untuk simulasinya kalian bisa buat table dengan struktur seperti ini ya

+-----------+-----------------------+------+-----+---------------------+----------------+
| Field     | Type                  | Null | Key | Default             | Extra          |
+-----------+-----------------------+------+-----+---------------------+----------------+
| id        | int(11)               | NO   | PRI | NULL                | auto_increment |
| username  | varchar(40)           | NO   |     | NULL                |                |
| email     | varchar(100)          | NO   |     | NULL                |                |
| gender    | enum('pria','wanita') | YES  |     | NULL                |                |
| password  | text                  | NO   |     | NULL                |                |
| create_at | timestamp             | NO   |     | current_timestamp() |                |
+-----------+-----------------------+------+-----+---------------------+----------------+


nah kal

*/

package main

import (
	"context"
	"database/sql"
	"fmt"
	"sql-injection/db"
	"time"
)

type Users struct {
	Id       int
	Username string
	Email    string
	Gender   sql.NullString
	Password string
	CreateAt time.Time
}

func getDataSqlInjection(mysqlConn *sql.DB, username string, password string) {

	var ctx = context.Background()
	var rows, err = mysqlConn.QueryContext(ctx, "SELECT * FROM users WHERE username = '"+username+"' and password = '"+password+"'")
	if err != nil {
		panic(rows)
	}
	defer rows.Close()

	var allUsers []Users = []Users{}

	for rows.Next() {

		var user Users
		var err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Gender, &user.Password, &user.CreateAt)
		if err != nil {
			panic(err)
		}

		allUsers = append(allUsers, user)
	}

	fmt.Println(allUsers)
}

func getDataSqlInjectionFix(mysqlConn *sql.DB, username string, password string) {

	var ctx = context.Background()
	var rows, err = mysqlConn.QueryContext(ctx, "SELECT * FROM users WHERE username = ? and password = ?", username, password) // sql injection fix
	if err != nil {
		panic(rows)
	}
	defer rows.Close()

	var allUsers []Users = []Users{}

	for rows.Next() {

		var user Users
		var err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Gender, &user.Password, &user.CreateAt)
		if err != nil {
			panic(err)
		}

		allUsers = append(allUsers, user)
	}

	fmt.Println(allUsers)
}

func main() {
	var mysqlConn = db.CreateConnection()
	defer mysqlConn.Close()
	getDataSqlInjection(mysqlConn, "hasan", "123")
	/*

		kalo kita jalankan hasilnya sudah benar ya,
		dan kalo kalian jalankan dengan password yang salah maka hasilnya adalah
		slice kosong ya


		sekarang kita coba tambahkan sql injectionnya
		apa itu sql injection ?

		sql injection adalah sebuah teknik yang menyalangunakan sebuah celah keamanan yang terjadi dalam
		lapisan baris data sebuah applikasi pada perintah sql

		biasanya sql injection dilakukan dengan mengirimkan input user dengan perintah yang salah
		sehingga menyebabkan hasil sql yang kita buat tidak valid


		akan tetapi dengan menambahkan saja sql injection dan ketika saya menuliskan passwordnya
		salah maka kita tetap saja berhasil mendapatkan datanya

		jadi ini sangat berbahaya, jika kita salah membuat sql, maka bisa jadi data kita tidak aman
		nah sekarang kita coba memasukkan perintah sql injectionnya
	*/

	getDataSqlInjection(mysqlConn, "hasan '; #", "salah")
	/*

		nah kalo kita jalankan datanya tetap terambil meskipun passwordnya salah
		inilah sql injection

		nah pagar ini artinya adalah komentar
		maka pengecheckan yang password akan dikomentari dan tidak dijalankan
		inilah kenapa hasilnya ada data hasan yang keambil meskipun passwordnya salah

		nah cara mengatasinya kalian bisa gunakan parameter pada
		funciton query ata execute

		tinggal kita masukkan inputan dari usernya pada parameter ke 3
		dan menambahkan tanda ? pada posisi disqlnya

		dan parameter ke 3 ini adalah variadic function
		artinya kalian bebas memasukkan berapa banyak parameter didalamnya

		jadi seperti ini

	*/

	getDataSqlInjectionFix(mysqlConn, "hasan '; #", "salah")

	/*

		jadi kalo kita jalankan maka akan mengembalikan slice kosong
		nah ini sudah benar ya


		jadi mau perintah sql apapun yang membutuhkan parameter seperti insert, delete, update ataupun select
		maka gunakanlah parameter agar menghindari sql injection

		jadi seperti itu ya sql injection dan cara mengatasinya
		mudah mudahan kalian paham

	*/
}
