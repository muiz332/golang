/*

- request body



dimateri kali ini kita akan belajar bagaimana cara mendapatkan data dari request
yang dikirimkan pada form

akan tetapi sebelum kalian mendapatkan datanya, kalian harus parse dulu
jadi parse itu agar datanya bisa terbaca

kalo kalian belajar expres untuk memparsenya itu sangat mudah hanya menuliskan

app.use(app.json()) dan dia akan memperse seluruh request

akan tetapi kalo digolang, kita harus memparse direquest yang membutuhkan data
bodynya

caranya kalian cukup menuliskan r.ParseForm()
dan dia mengembalikan error

dan untuk mendapatkan datanya kalian tulis

r.PostForm.Get(keynya)

oke langsung saja kita coba


*/

package request_body

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)




func handleBody(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm() // parse dulu

	if err != nil{
		panic(err)
	}

	data := r.PostForm 

	firstName := data.Get("firstname") // ambil datanya
	lastName := data.Get("lastname")

	// r.PostFormValue("firstname") 

	/*
	
	nah method ini r.PostFormValue("firstname") digunakan untuk mengambil datanya 
	tanpa perlu secara manual memparsingnya

	jadi sebenarnya dibelakang layar pada method PostFormValue itu juga memparsing 
	terlebih dahulu
	
	*/


	fmt.Fprintf(w,"selamat datang %s%s",firstName,lastName)
}

func TestFormBody(t *testing.T) {

	
	reqBody := strings.NewReader("firstname=muiz&lastname=kuy") // mengembalikan object reader

	/*
	
	jadi kita buat request bodynya menggunakan NewReader yang nanti kita kirimkan
	pada parameter ke 3 pada function NewRequest

	jadi kalo kalian perhatikan mirip seperti query string ya
	memang sebenarnya seperti itu spesifikasinya

	bedanya kalo post itu dikirimkan bukan lewat url tetapi lewat form
	
	*/
	req := httptest.NewRequest(http.MethodPost,"http://localhost:8080/",reqBody)
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")

	/*
	
	dan jangan lupa kalian harus mengirimkan content typenya juga
	dengan cara set headernya 


	
	*/

	res := httptest.NewRecorder()

	handleBody(res,req)

	result := res.Result()
	defer result.Body.Close()
	
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t,"selamat datang muizkuy",string(body))
	
}