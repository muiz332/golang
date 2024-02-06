/*


- query parameter


dimateri kali ini kita akan membahas tentang query parameter
lebih tepatnya adalah query string ya

jadi tanda tanya ? setelah urlnya dan bagaimana golang menangkap
data yang dikirimkan lewat parameter ?

oke langsung saja kita praktek



*/

package query_parameter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"query_parameter/routers"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var url string = ""

func TestMain(m *testing.M) {

	server := httptest.NewServer(routers.Router())
	defer server.Close()

	url = server.URL
	m.Run()
}

func TestWitoutQuery(t *testing.T) {
	res, err := http.Get(url + "/home")
	if err != nil {
		require.Error(t, err)
	}

	body := res.Body
	defer body.Close()

	readBody, err := io.ReadAll(body)
	if err != nil {
		require.Error(t, err)
	}

	require.Equal(t, "hello", string(readBody))

}

func TestWithQuery(t *testing.T) {
	res, err := http.Get(url + "/home?nama=muiz")
	/*

		jadi kita tinggal tambahkan seperti ini untuk menambahkan
		query string ke url

	*/
	if err != nil {
		require.Error(t, err)
	}

	body := res.Body
	defer body.Close()

	readBody, err := io.ReadAll(body)
	if err != nil {
		require.Error(t, err)
	}

	assert.Equal(t, "hello muiz", string(readBody))

}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]

	data := strings.Join(names, ",")

	w.Write([]byte(data))

}

func TestMultipleQuery(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=muiz&name=hasan", nil)
	recorder := httptest.NewRecorder()

	HandleAbout(recorder, req)

	res := recorder.Result()
	body := res.Body
	defer body.Close()

	result, _ := io.ReadAll(body)

	assert.Equal(t, "muiz,hasan", string(result))

	/*

		jadi kita bisa tangkap multiple query dengan cara kita memasukkan keynya itu sama
		jadi sebenarnya url.Query itu mengembalikan map

		dan tenang saja ketika memiliki key yang sama itu tidak akan mereplacenya
		karena sebenarnya isi dari mapnya itu adalah slice



	*/
}
