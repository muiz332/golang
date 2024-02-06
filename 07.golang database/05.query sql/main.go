/*

-  query sql

untuk operasi sql yang tidak membutuhkan hasil kita bisa gunakan exac, akan tetapi
kalo kalian butuh hasil

misalkan kalian ingin menampilkan data dari database, maka kalian butuh yang namanya
querycontext

oke langsung saja ya kita coba



*/

package main

import (
	"context"
	"database/sql"
	"fmt"
	"query-sql/db"
)

type Orang struct {
	Id     int
	Nama   string
	Alamat string
}

func main() {

	var mysqlConn = db.CreateConnection()
	defer mysqlConn.Close()

	var ctx = context.Background()
	var raws, err = mysqlConn.QueryContext(ctx, "SELECT * FROM orang")
	if err != nil {
		panic(err)
	}

	defer raws.Close() // harus diclose

	var data []Orang = []Orang{}

	for raws.Next() {

		var result Orang
		var err = raws.Scan(&result.Id, &result.Nama, &result.Alamat) // assigment ke var result sesuai urutan

		if err != nil {
			panic(err)

			/*

				error terjadi kalo misalkan kalian memasukkan type data pada method Scan itu
				tidak sesuai dengan type data yang ada dicoloumn pada table yang kalian buat

			*/
		}

		data = append(data, result) // setiap looping akan menambahkan data ke slice

		/*

			method Next() ketika dipanggil dan ada datanya didalam database maka returnnya true
			kalo tidak ada maka false

			jadi kita looping dengan kondisinya next(), artinya ketika datanya masih ada
			maka looping akan tetap dijalankan

			method Scan() digunakan untuk mendapatkan data dari database, dan data tersebut
			dapat kita masukkan kedalam variable biasa atau struct

			perlu kalian ingat bahwa yang kalian masukkan ke kedalam parameter scan harus berupa
			pointer, yang artinya harus menambahkan tanda & sebelum nama variablenya

			jadi setiap looping kita tampung datanya ke slice, kemudian setelah semua datanya
			ada didalam slice, barulah kita tampilkan datanya diterminal

		*/
	}

	fmt.Println(data)

	/*

		oke jadi seperti ini ya cara query ke database yang membutuhkan hasil data, kalian bisa
		gunakan queryContext

		dan untuk type data yang dibutuhkan pada saat pemanggilan method scan
		itu harus sesuai dengan column yang ada didalam database ya

		jadi kalian harus tulis seperti ini

		column ditable					type data
		varchar, char 					string
		int, bigint						int32,int64,int
		float,double					float32, float64
		boolean							boolean
		date,datetime,time, timestamp   time.time


		akan tetapi kalo type data date,datetme,time,timestamp itu sebanarnya database tidak mengembalikan type data
		time.time akan tetapi mengembalikan []byte, dan ini bisa dikonversi menjadi string, dan dari string
		dikonversi ke type data type

		tetapi ini agak ribet ya, nah untungnya driver mysql sudah menyediakan auto parse untuk
		mengatasi hal tersebut

		dengan cara kalian tuliskan saja pada data sourcenya
		?parseTime=true

		root:@tcp(localhost:3306)/test?parseTime=true

		jadi cukup seperti itu saja maka error yang terjadi karena kesalahan type data
		akan berhasil ditangani

		akan tetapi kalo didalam database kalian ada column yang nullable
		artinya datanya itu boleh diisi atau tidak, kalo tidak diisi biasanya kita bisa set null didatabasenya

		nah golang tidak mengerti type data null ini, karena kan type data seperti string, time.time, int boolean float
		itukan tidak bisa kita kasih nilai nill

		nah lalu bagaimana cara mengatasinya?
		jadi kalo misalkan ada column yang nullable kalian bisa tulis seperti ini
	*/

	type TestTypeData struct {
		Id     int
		Nama   string
		Umur   sql.NullInt16
		Alamat sql.NullString
	}

	/*
		karena driver mysql itu tidak mendukung hal tersebut
		jadi kalian harus gunakan type data yang telah disediakan oleh package sql
		seperti itu ya

		lalu dia akan mengembalikan struct yang isinya seperti ini

		type NullInt16 struct{
			Int16 int16
			Valid bool
		}

		nah kita bisa check nih kalo misalkan valid itu true maka datanya ada
		kalo misalkan valid itu false maka datanya tidak ada atau nill

		oke jadi seperti itu ya bagaimana cara kita mendapatkan data dari
		database

		mudah mudahan kalian paham




	*/

}
