/*

- database pooling


saat kita membuat sql.Db object digolang sebenarnya bukanlah sebuah connection kedatabase
melainkan sebuah pool ke database

jadi saat kalian membuat open database itu kalian sebenarnya membuat pool kedatabase
jadi konsepnya adalah database pooling

database pooling itu adalah kumpulan dari connection, didalam sql.Db golang itu melakukan management koneksi database  secara
otomatis, hal ini menjadikan kita tidak perlu melakukan management koneksi secara manual

jadi secara default golang sudah menerapkan kemampuan database pooling, dimana itu sangat bagus ya
lebih canggih karena disaat bahasa pemrograman yang lain membutuhkan library tambahkan

kalo digolang sudah secara default menggunakan kemampuan database pooling ini
dengan kempuan database pooling ini, kita bisa menentukan jumlah minimal dan maximal
koneksi yang dibuat oleh golang

sehingga tidak membanjiri koneksi ke database, karena biasanya ada batas maximal koneksi
yang bisa ditangani oleh database yang kita gunakan

misalkan nanti kalo ada 1 user yang melakukan request dan kita tidak batasi koneksi nya
maka golang akan membuat 1 koneksi baru

bayangkan kalo ada 1000 user yang melakukan request, maka golang juga akan membuat 1000
koneksi, dan ingat database itu memiliki limit koneksi yang kecil

misalkan kalo dimysql itu limitnya sekitar 151 koneksi
jadi itulah mengapa kita harus membatasi koneksi ke database kita

akan tetapi kita harus mengatur minimal koneksi, karena pada saat applikasi kita jalankan
applikasi kita sudah siap dengan minimal koneksi yang telah dibuat

jadi secara default golang akan membuat 10 minimal koneksi ke database
dan misalkan kalian mengatur maksimal 50 koneksi

ketika ada 1000 user yang request, maka sisanya akan mengantri terlebih dahulu
menunggu 50 koneksi tersebut selesai menjalankan tugasnya



- method yang digunakan untuk mengatur database pooling


db.SetMaxIdleConns(number) => untuk mengatur minimal koneksi yang dibuat diawal
db.SetMaxOpenConns(number) => mengatur maximal koneksi yang dibuat


db.SetMaxIdleTime(duration)
	pengaturan berapa lama koneksi baru yang sudah tidak digunakan
	itu dihapus

	jadi kalo minimal 10 koneksi dan ada 5 koneksi baru, maka setelah 5 koneksi itu digunakan
	dan tidak digunakan lagi, maka berapa lama 5 koneksi itu dihapus

db.SetConnsMaxLifeTIme(duration)
	pengaturan berapa lama koneksi boleh digunakan, dan tidak mengenal idle connection atau open connection
	nah kalo misalkan kita set 1 jam, maka akan diclose koneksinya

	nah kalo koneksi yang minimalnya, maka dia akan diclose akan diganti dengan
	koneksi yang baru atau ini istilahnya memperbaharui koneksi yang sudah lama

jadi ini pengaturan database pooling digolang
oke kita coba

*/

package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func main() {

	mysqlConn := CreateConnection()
	defer mysqlConn.Close()
	fmt.Println("database berhasil terkoneksi")
}
