package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiber(t *testing.T) {


	app := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 5,
		ReadTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
	})


	app.Get("/",func(c *fiber.Ctx) error {

		nama := c.Query("nama")
		fmt.Println(nama)

		cookie := c.Cookies("login")
		fmt.Println(cookie)


		return c.SendString("hello")
	})


	// cara mengetest

	req := httptest.NewRequest(http.MethodGet,"/?nama=hasan",nil)

	cookie := new(http.Cookie)
	cookie.Name = "login"
	cookie.Value = "sudah login"


	// req.AddCookie(cookie)

	req.AddCookie(&http.Cookie{
		Name: "login",
		Value: "sudah login",
	})

	res, err := app.Test(req)

	assert.Nil(t,err)

	body := res.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	assert.Nil(t,err)

	assert.Equal(t,"hello", string(data))
	assert.Equal(t,200,res.StatusCode)


}