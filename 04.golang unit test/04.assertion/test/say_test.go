package test

import (
	"assertion/helper"
	"fmt"
	"testing"

	"github.com/Muizzuddin-github/sayhello"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHello(t *testing.T) {

	var result string = helper.SayHello("muiz")
	var expect string = "hello muiz"
	assert.Equal(t, expect, result, fmt.Sprintf("expect %s result %s", expect, result))
	fmt.Println("selesai dijalankan")

	/*

		jadi assert ini dibelakang layar  akan menjalankan t.Fail()

		jadi program yang ada dibawahnya akan tetap dijalankan

		misalkan kita ingin mengetest menggunakan function assert.Equal() ada 4 buah parameter

		parameter pertama itu struct dari pointer testing.T
		parameter kedua itu nilai yang kita harapkan atau expectation
		parameter ketiga itu hasil dari code yang kita test
		parameter keempat itu adalah keterangan ketika testnya gagal (optional)

		kita jalankan

		go test -v ./test

		maka kan muncul ya hasilnya

	*/
}

func TestSayHello2(t *testing.T) {
	var result string = helper.SayHello("hasa")
	var expect string = "hello hasan"
	require.Equal(t, expect, result)

	/*

		kalo kita menggunakan require, maka yang dijalankan adalah t.FailNow()

		maka code program yang dibawahnya tidak akan dijalankan ketika terjadi error
		dan ketika menggunakan testify informati errornya itu akan lebih detil



	*/

	result = sayhello.SayHello()

}
