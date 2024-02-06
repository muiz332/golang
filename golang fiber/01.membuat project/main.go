/*

- golang fiber

jadi golang fiber ini adalah framwork yang memudahkan kalian untuk membuat
server menggunakan golang


fiber terinspirasi dari express, jadi nanti struktur codenya itu mirip
dengan express

oke langsung saja kita coba



*/

package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main(){

	app := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 5, // batas waktu koneksi aktif
		ReadTimeout: time.Second * 5, // request timeout
		WriteTimeout: time.Second* 5, // response time out
	})

	app.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})


	app.Listen("localhost:8080")

}