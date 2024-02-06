/*

- response status code

jadi kita perlu mengirimkan status code ya pada response yang kita
berikan

kalian bisa baca disini untuk http status codenya

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status

secara default response code yang diterima itu 200
kalo kalian ingin mengubahnya kalian bisa tuliskan di

r.WriteHeader(int)

kita coba


*/

package request_body

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)




func handleGet(w http.ResponseWriter, r *http.Request){

	nama := r.URL.Query().Get("nama")

	if nama == ""{
		w.WriteHeader(400)
		fmt.Fprintf(w,"bad request")
	}else{
		w.WriteHeader(200)
		fmt.Fprintf(w,"berhasil")
	}

}

func TestStatusCode(t *testing.T) {

	
	req := httptest.NewRequest(http.MethodPost,"http://localhost:8080/",nil)
	res := httptest.NewRecorder()

	handleGet(res,req)

	result := res.Result()
	defer result.Body.Close()

	assert.Equal(t,400,result.StatusCode)
	assert.NotEqual(t,200,result.StatusCode)

	
}