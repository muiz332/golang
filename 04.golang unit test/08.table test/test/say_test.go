package test

import (
	"table-test/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSay(t *testing.T) {

	type Person struct {
		Nama        string
		Param       string
		Expectation string
	}

	var data []Person = []Person{
		{
			Nama:        "muiz",
			Param:       "muiz",
			Expectation: "hello muiz",
		},
		{
			Nama:        "hasan",
			Param:       "hasan",
			Expectation: "hello hasan",
		},
		{
			Nama:        "husain",
			Param:       "husain",
			Expectation: "hello husain",
		},
	}

	for _, p := range data {
		t.Run(p.Nama, func(t *testing.T) {
			var result string = helper.SayHello(p.Param)
			require.Equal(t, p.Expectation, result)
		})
	}
}

/*

kalo saya jalankan

=== RUN   TestSay
=== RUN   TestSay/muiz
=== RUN   TestSay/hasan
=== RUN   TestSay/husain
--- PASS: TestSay (0.00s)
    --- PASS: TestSay/muiz (0.00s)
    --- PASS: TestSay/hasan (0.00s)
    --- PASS: TestSay/husain (0.00s)


maka semua data yang ada didalam struct berhasil kita test
jadi kalo kalian punya sub test dan masing masing sub test itu yang berbeda hanya
parameternya

maka kalian bisa gunakan konsep table test, jadi kalo ingin menambah parameternya
tinggal kalian tambahkan data dislicenya ya

jadi itulah konsep table test
mudah muahan kalian paham

*/
