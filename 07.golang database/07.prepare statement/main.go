/*

- prepare statement

saat kita menggunakan function query atau exac yang menggunakan parameter, sebenarnya implementasi
dibawah menggunakan prepare statement

jadi prepare statement ini adalah koneksi yang terikat, jadi sebelum kalian menggunakan query atau exac
maka golang akan menyiapkan koneksi yang diambil dari connection pooling

kadang ada kasus kita ingin melakukan beberapa hal yang sama sekali gus hanya berbeda
parameternya saja

jadi perintah sqlnya sama tetapi hanya parameternya saja yang berbeda
misalkan kalian ingin melakukan insert many ke database

nah itukan yang berbeda hanya parameternya saja, jadi kalian menggunakan exac dan ingin
melakukan insert lebih dari satu maka yang terjadi adalah

dibelakang layar prepare statement disiapkan oleh golang, nah masalahnya kalo misalkan kalian bikin
exac atau query lebih dari satu tetapi perintah sqlnya sama

maka akan dibuatkan prepare statement lebih dari satu juga sesuai pemanggilan query atau exacnya
nah cara kerja dari prepare statement adalah

golang mengecheck apakah ada connection yang kosong dipoolnya, kalo ada maka diambil
dan digunakan untuk berinteraksi dengan database, setelah itu dikembalikan lagi

kalo misalkan kita menjalankan query atau exac lebih dari satu dengan sintax sql yang sama
dan hanya parameternya saja yang berbeda

maka itu akan menjadi lambar karena setiap kita ingin melakukan query atau exac akan bertanya
dulu ke poolnya

bagaimana kalo misalkan kita buat prepare statement sendiri
dan prepare statement tersebut yang isinya koneksi akan digunakan untuk berinteraksi
dengan database hanya menggunakan 1 koneksi tersebut yang telah diambil dari connection pool

jadi dengan 1 connection kita bisa jalankan perintah query atau exac yang sama
dan itu akan lebih cepat

tanpa harus mengembalikan connection nya sebelum semua perintah insert dijalankan

jadi nanti kita buat prepare statement 1 kali sebelum menjalankan banyak insert
oke langsung saja kita coba





*/

package main

import (
	"context"
	"fmt"
	"prepare-statement/db"
	"strconv"
)

func main() {

	var mysqlConn = db.CreateConnection()
	defer mysqlConn.Close()

	var ctx context.Context = context.Background()
	var sintaxSql string = "INSERT INTO users (username,email,password) VALUES (?,?,?)"
	var pre, err = mysqlConn.PrepareContext(ctx, sintaxSql) // prepare connection
	if err != nil {
		panic(err)
	}
	defer pre.Close() // prepare connection harus diclose

	for i := 0; i < 10; i++ {
		var username string = "username" + strconv.Itoa(i)
		var email string = "email" + strconv.Itoa(i) + "@gmail.com"
		var password string = "password" + strconv.Itoa(i)

		var result, err = pre.ExecContext(ctx, username, email, password)
		/*

			execContext dalam context yang sama dengan prepare connectionn6ya
			dan didalam parameternya kita hanya memasukkan parameternya saja

			karena sintax sql nya sudah ada didalam prepare connectionnya

		*/
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("insert id", id)

	}

	/*

		jadi seperti itu ya, kalo kalian ingin mengekesekusi banyak perintah yang sama tetapi hanya
		parameternya saja yang berbeda, misalkan kalian melakukan insert many

		nah kalian harus menggunakan prepare connection agar ketika insert program kalian bisa lebih
		cepat dan tidak lambat

		jadi seperti itu tentang prepare connection
		mudah mudahan kalian paham

	*/
}
