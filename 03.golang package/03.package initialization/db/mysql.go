package db

import "fmt"

var conn string = ""

func init() {
	conn = "mysql"
	fmt.Println("mysql terkoneksi")
}

func GetConn() {
	fmt.Println(conn)
}
