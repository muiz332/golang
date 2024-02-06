/*

- request form



*/

package golang_fiber

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
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


// go:embed source/contoh.txt
var contohFIle []byte 

func TestRequestForm(t *testing.T){


	app.Post("/upload",func(c *fiber.Ctx) error {

		file, err := c.FormFile("file") // untuk mendapatkan data file
		if err != nil{
			return err
		}

		// untuk menjadikan file sebagai file asli

		err = c.SaveFile(file,"./target/" + file.Filename)
		if err != nil{
			return err
		}


		return c.SendString("file berhasil diupload")
	})


	// buat request body

	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	file, _ := writer.CreateFormFile("file","contoh.txt")
	file.Write(contohFIle)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost,"/upload",reqBody)
	req.Header.Set("Content-Type",writer.FormDataContentType())
	res, err := app.Test(req)

	assert.Nil(t,err)

	body := res.Body
	defer body.Close()

	dataByte, err := io.ReadAll(body)
	assert.Nil(t,err)

	assert.Equal(t,"file berhasil diupload",string(dataByte))
	assert.Equal(t,200,res.StatusCode)





}

