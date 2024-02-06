package request_body

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Custom config
var app = fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName: "Test App v1.0.1",
	IdleTimeout: time.Second * 5,
	ReadTimeout: time.Second * 5,
	WriteTimeout: time.Second * 5,
})


type Body struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func TestRequestBody (t *testing.T){
	app.Post("/login",func(c *fiber.Ctx) error {

	
		body := Body{}

		err := json.Unmarshal(c.Body(),&body)
		if err != nil{
			c.SendStatus(fiber.StatusBadRequest)
			return c.SendString(err.Error())
		}

		c.SendStatus(fiber.StatusOK)
		return c.SendString(fmt.Sprintf("username %s",body.Username))

	})

	// request body
	bodySend := strings.NewReader(`{
		"username" : "muiz",
		"email" : "muiz@gmail.com",
		"password" : "1234"
	}`)
	req := httptest.NewRequest(http.MethodPost,"/login",bodySend)
	req.Header.Set("Content-Type","application/json")


	res, err := app.Test(req)
	assert.Nil(t,err)

	body := res.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	assert.Nil(t,err)

	assert.Equal(t,200,res.StatusCode)
	assert.Equal(t,"username muiz",string(data))

	fmt.Println(string(data))

	/*
	
	jadi kalian bisa gunakan c.Body untuk menangkap data jsonnya dan dari situ kita bisa
	encoding dari json ke struct menggunakan function unmarshal

	
	*/
}


type Register struct{
	Username string `json:"username" form:"username"` 
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}


func TestRegister(t *testing.T) {
	app.Post("/register",func(c *fiber.Ctx) error {

	
		body := Register{}

		err := c.BodyParser(&body)
		if err != nil{
			c.SendStatus(fiber.StatusBadRequest)
			return c.SendString(err.Error())
		}

		c.SendStatus(fiber.StatusOK)
		return c.SendString(fmt.Sprintf("username %s",body.Username))

	})

	// request body
	bodySend := strings.NewReader(`{
		"username" : "muiz",
		"email" : "muiz@gmail.com",
		"password" : "1234"
	}`)
	req := httptest.NewRequest(http.MethodPost,"/login",bodySend)
	req.Header.Set("Content-Type","application/json")


	res, err := app.Test(req)
	assert.Nil(t,err)

	body := res.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	assert.Nil(t,err)

	assert.Equal(t,200,res.StatusCode)
	assert.Equal(t,"username muiz",string(data))

	fmt.Println(string(data))

	/*
	
	fiber sebenarnya sudah menyediakan function yaitu body parser untuk memparsing data
	yang dikirimkan ke structnya tanpa melakukan konversi secara manual seperti tadi

	dengan menambahkan tag json atau form kita bisa melakukan konversi data yang dikirimkan
	ke struct secara otomatis

	jadi seperti itu ya mudah mudahan kalian paham
	
	*/
}