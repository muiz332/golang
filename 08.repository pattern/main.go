/*

- repository pattern


kali ini kita akan bahas bagaimana cara untuk membuat applikasi berbasis database
dengan struktur folder yang baik

nah kalo kalian tahu struktur folder pada sebuah applikasi yang menerapkan MVC
atau Model View Controller

pada structor folder tersebut terdapat kelemahan pada folder controller
dimana isi folder controller adalah busines logic dan aksi yang berhubungan dengan
database

seperti insert, update dan lain lain

nah repository pattern itu sebenarnya memecah bagian dari folder controller menjadi dua
yang ada didalam folder controller adalah busines logicnya dan interfaksi terhadap
databasenya itu ada direpositorynya

akan tetapi seperti yang kita tahu golang itu adalah bahasa pemrograman yang static type
dimana harus mendefinisikan dulu type datanya

nah kita buat dulu repository interface, dimana didalamnya hanya terdapat interface yang
digunakan untuk kontak pada struct yang ada difolder repository implement

folder repository implement isinya hanya sebuah struct yang memiliki method
untuk mengakses ke database

selanjutnya ada folder entity, dimana ini representasi type data dari table table
yang ada didalam database

jadi seperti itulah struktur folder yang baik dengan menggunakan pendekatan repository
pattern

mudah mudahan kalian paham


*/

package main

import (
	"fmt"
	"repository_pattern/controllers"
	"repository_pattern/db"
)

func main() {
	mysqConn := db.CreateConnection()
	defer mysqConn.Close()

	result, err := controllers.GetAllUsers(mysqConn)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}
