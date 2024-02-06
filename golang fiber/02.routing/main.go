/*

- routing



*/

package main

import (
	"fmt"
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

	
	app.Get("/mahasiswa/:id",func(c *fiber.Ctx) error {
		
		fmt.Println(c.Params("id"))

		// atau bisa menggunakan c.AllParams()



		return c.SendString("mahasiswa")
	})


	app.Listen("localhost:8080")

	/*
	
	jadi kalo kalian sudah belajar express maka ini sangat direkomendasikan
	menggunakan framework fiber

	akan tetapi yang berbeda adalah request dan response itu menjadi satu
	pada callbacknya 

	
	
	*/

}