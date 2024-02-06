package tests

import (
	"beforeAll-afterAll/helper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {

	fmt.Println("sebelum unit test dijalankan")

	m.Run() // menjalankan semua unit test

	fmt.Println("setelah unit test dijalankan")

}

func TestSayName(t *testing.T) {
	var result string = helper.SayName("muiz")
	require.Equal(t, "hello muiz", result)
}

func TestSayName2(t *testing.T) {
	var result string = helper.SayName("hasan")
	require.Equal(t, "hello hasan", result)
}

/*

kalo kita jalankan maka akan muncul

sebelum unit test dijalankan                  ter
=== RUN   TestSayName
--- PASS: TestSayName (0.00s)
=== RUN   TestSayName2                        ow\
--- PASS: TestSayName2 (0.00s)                ter
PASS
setelah unit test dijalankan

jadi before dan after all ini biasanya digunakan untuk mengatur koneksi ke database
jadi beforenya kita akan koneksikan kedatabase dan afternya kita akan tutup koneksi kedatabase

jadi seperti itu ya mudah mudahan kalian paham

*/
