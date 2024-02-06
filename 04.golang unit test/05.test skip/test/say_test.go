package test

import (
	"test-skip/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSayHello(t *testing.T) {

	t.Skip("skip unit test")
	var result string = helper.SayHello("muiz")
	var exp string = "hello muiz"
	require.Equal(t, exp, result)

}
func TestSayHello2(t *testing.T) {

	var result string = helper.SayHello("hasan")
	var exp string = "hello hasan"
	require.Equal(t, exp, result)
}

/*


jadi kalian cukup gunakan function t.Skip() untuk melakukan skipping atau melewati
function unit testnya

kalo kita jalankan maka munculnya akan seperti ini

=== RUN   TestSayHello
    say_test.go:12: skip unit test
--- SKIP: TestSayHello (0.00s)
=== RUN   TestSayHello2
--- PASS: TestSayHello2 (0.00s)
PASS
ok      test-skip/test  (cached)

ada tulisannya skip pada function TestSayhello
dan diatasnya ada tuliskan file unit test mana beserta baris berapa skip nya

oke jadi itulah cara melakukan skip unit test digolang
mudah muahan kalian paham

*/
