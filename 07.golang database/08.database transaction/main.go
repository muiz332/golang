/*

- database transaction

salah satu fitur andalan didatabase adalah database transaction
golang juga bisa menggunakan fitur database transaction ini

jadi kalian harus mengerti dulu ya bagaimana cara kerja database transaction
dimysql

jadi kali ini kita hanya fokus bagaimana cara menggunakan fitur database
transaction digolang

jadi secara default semua perintah sql yang kita kirim menggunakan golang akan secara
otomatis dicommit atau langsung dimasukkan kedalam database

namun jika kita menggunakan fitur transaction, perintah sql yang kita kirim tidak
secara langsung akan masuk kedalam database kalo kita tidak secara manual
menuliskan commit

nah kanapa kita perlu belajar database transaction ini?

saat membuat applikasi berbasis database, jarang sekali kita menggunakan satu jenis perintah
sql per aksi yang digunakan

contoh ketika customer memesan tombol pesan ada banyak yang harus kita lakukan

- memasukkan data pesanan ke order
- membuat data detail pesanan ditable order detail
- menurunkan quantity pada table product
- dan lain lain

artinya bisa saja dalam satu aksi kita bisa melakukan beberapan perintah sekaligus
jika terjadi kesalahan disalah satu perintahnya, harapannya adalah perintah perintah
sebelumnya dibatalkan, agar data tetap konsisten

jadi database transaction ini adalah sebuah fitur dimana kita bisa menggunakan beberapa
perintah sql menjadi satu kesatuan

jadi ketika ada salah satu saja perintah sql yang gagal maka semuanya akan
dirollback atau dikembalikan ke kondisi awal


nah untuk menggunakan fitur transaction digolang, kita cukup menggunakan function begin
dimana dia akan menghasilkan sebuah struct tx yang merupakan representasi dari
transaction

nah struct ini yang nanti akan kita gunakan sebagai pengganti dari struct db
dimana hampir semua transaction di db juga ada di tx

setelah selesai proses transaksi, kita bisa menggunakan function commit atau rollback
function rollback digunakan untuk mambatalkan semua perintah sebelumnya
agar tidak masuk kedalam database

nah function begin akan secara otomatis menggunakan koneksi dari pool yang nanti
satu koneksi tersebut bisa digunakan oleh beberapa query




*/

package main

import (
	"context"
	"database-transaction/db"
)

func main() {

	var mysqlConn = db.CreateConnection()
	defer mysqlConn.Close()

	tx, err := mysqlConn.Begin()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	result, err := tx.ExecContext(ctx, "INSERT INTO users (username,email,gender,password) VALUES (?,?,?,?)", "faix", "faix@gmail.com", "pria", "1234")
	if err != nil {
		tx.Rollback() // kalo error kita rolback
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	if _, err := tx.ExecContext(ctx, "UPDATE users SET username = ? WHERE id = ?", "faix kuy", id); err != nil {
		tx.Rollback() // kalo error kita rolback
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	/*

		maka kalo kita jalankan hasilnya seperti ini
		 14 | romi kuy  | romi@gmail.com      | pria   | 1234      | 2023-11-21 20:49:14

		 akan tetapi kalo terjadi error maka tidak akan masuk kedatabase
		 jadi itulah database transaction

		 kalian sebaiknya gunakan database transaction karena pada applikasi berbasis
		 database itu ketika kita bikin sebuah fitur

		 query sqlnya kebanyakan lebih dari satu, dan ketika ada yang perlu dirolback
		 maka gunakanlah database transaction

		 jadi itu ya mengenai database transaction
		 mudah mudahan kalian paham

	*/

}
