/*

- request form



*/

package golang_fiber

import (
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


var app = fiber.New(fiber.Config{
	IdleTimeout: time.Second * 5,
	ReadTimeout: time.Second * 5,
	WriteTimeout: time.Second * 5,
})

func TestRequestForm(t *testing.T){


	app.Get("/",func(c *fiber.Ctx) error {


		// mendapatkan request body
		username := c.FormValue("username")
		email := c.FormValue("email") 

		return c.SendString(fmt.Sprintf("hallo username saya %s dan email saya %s",username,email))
	})

	reqBody := strings.NewReader("username=muiz&email=muiz@gmail.com")
	req := httptest.NewRequest(http.MethodGet,"/",reqBody)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	res, err := app.Test(req)

	assert.Nil(t,err)

	body := res.Body
	defer body.Close()

	dataByte, err := io.ReadAll(body)
	assert.Nil(t,err)

	assert.Equal(t,"hallo username saya muiz dan email saya muiz@gmail.com",string(dataByte))
	assert.Equal(t,200,res.StatusCode)


}

