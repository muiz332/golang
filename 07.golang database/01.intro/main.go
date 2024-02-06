/*


- package database



sekarang kita akan mengenal dulu mengenai package database digolang
jadi bahasa pemrograman golang itu secara default memiliki sebuah package bernama database

package database ini adalah package yang berisikan kumpulan standard interface yang menjadi
standard untuk berkomunikasi ke database

nah keuntungannya apa dengan adanya package standard database?
jadi kalo kalian ingin mengakses database apapun itu codenya sama

jadi nanti kalian ingin mengakses databasenya mysql, sqlite, postgre, oracel
itu semuanya kalian tidak perlu ganti code golang semuanya sudah standard
mengikuti standard yang sama sesuai dengan package database

mungkin yang membedakan hanya code sqlnya saja, jadi ini sangat bagus sekali ya
dengan adanya package database kita tidak perlu mengganti code kita lagi
ketika kita ingin migrate database



- cara kerja package database

applikasi => database interface => database driver => DBMS

jadi applikasi yang akan kita buat nanti memanggil database inteface, jadi nanti akan memanggil
interface interface atau struct yang ada didalam package database

nah tapi sebelumnya itu tidak secara otomatis langsung bisa digunakan, nah kalo kalian ingin
terkoneksi ke DBMSnya kalian tetap membutuhkan drivernya

jadi nanti dari applikasi kita dia akan memanggil database interfce dari database interface
dia akan memanggil database driver, barulah dari database driver akan memanggil DBMSnya

nah yang kita lakukan hanyalah menginstall database drivernya sesuai dengan database apa
yang nanti kita gunakna, pada materi ini nanti kita akan gunakan mysql

setelah install database driver menggunakan go get, nanti kalian hanya berinteraksi
dengan database interfacenya saja



- menambah database driver

jadi sebelum kita membuat program kita harus install dulu database drivernya
tanpa driver maka package database tidak mengerti apapun karena hanya berisi kontrak
interface interface saja

untuk menginstall driver kalian bisa tulis

go get -u github.com/go-sql-driver/mysql



setelah itu kalian tinggal import saja seperti ini

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



jadi kalian sudah bisa menginstall dan mengimport driver mysqlnya seperti itu

oke mudah mudahan kalian paham






*/

package main

func main() {

}
