package test

import (
	"fmt"
	"sub-test/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("before")
	m.Run()
	fmt.Println("after")
}

func TestSayHello(t *testing.T) {
	t.Run("hasan", func(t *testing.T) {
		var result string = helper.SayName("hasan")

		require.Equal(t, "hello hasan", result)
	})
	t.Run("say hello with parameter (muiz)", func(t *testing.T) {
		var result string = helper.SayName("muiz")

		require.Equal(t, "hello muiz", result)
	})
}

/*

jadi seperti itu sub test kita bisa membuat unit test didalam unit test
dan ini direkomendasikan ketika kalian ingin mengetest sebuah function

dari para kalian membuat banyak unit test lebih baik kalian membuat banyak sub test
karena itu bisa dikelompokkan

kalo kalian ingin mengetest function yang lain barulah kalian buat unit testnya

kita tinggal gunakan t.Run() dengan 2 buah parameter

parameter pertama itu adalah nama sub testnya dan parameter ke 2 itu adalah
function yang didalamnya kita bisa melakukan testing

untuk menjalankan sub unit test tertentu kalian cukup tuliskan

go test -v ./test -run "TestSayHello/hasan"

jadi itulah sub test
mudah muahan kalian paham

*/
