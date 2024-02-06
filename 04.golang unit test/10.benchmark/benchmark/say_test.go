package banchmark

import (
	"benchmark/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSay(t *testing.T) {
	var result string = helper.Say("muiz")
	require.Equal(t, "hello muiz", result)
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.Say("muiz")
	}

	/*

		nah untuk menjalankannya kalian tulis seperti ini

		go test -v -bench . nama_project_gomod/nama_foldernya

		hasilnya akan seperti ini

		=== RUN   TestSay
		--- PASS: TestSay (0.00s)
		goos: windows
		goarch: amd64
		pkg: benchmark/benchmark2
		cpu: Intel(R) Core(TM) i3-6006U CPU @ 2.00GHz
		BenchmarkSayHello
		BenchmarkSayHello-4     36677991                32.72 ns/op
		PASS
		ok      benchmark/banchmark2    2.368s


		akan tetapi kalo kalian lihat, unit test TestSay itu akan tetap dijalankan
		padahal saya hanya ingin menjalankan yang benchmarknya saja

		kalian bisa tulis seperti ini

		go test -v -run=$ -bench . nama_project_gomod/nama_foldernya

		jadi kalian tinggal tambahkan -run=$ agar unit testnya tidak akan dijalankan
		sebenarnya kalian tidak harus menuliskan tanda $ ya

		tetapi kalian bisa tulis kata apapun asalkan itu tidak sama dengan nama function
		unit test yang sudah kita buat

		karena golang itu akan menjalankan unit test sesuai namanya yang ada di -run=

		ketika kalian menuliskan namanya setelah tandah = di -run=
		dan nama tersebut tidak ada difunciton unit testnya

		maka golang tidak akan menjalankan unit testnya


	*/
}

func BenchmarkSub(b *testing.B) {

	b.Run("hasan", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			helper.Say("hasan")
		}
	})

	b.Run("selamat datang", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			helper.Say("selamat datang")
		}
	})

	/*

		jadi kita juga dapat membuat sub benchmark yang sama seperti unit test
		kita tetap menggunakan testing.B akan tetapi kita menggunakan funciton Run

	*/
}

func BenchmarkTable(b *testing.B) {

	type Says struct {
		Name   string
		Params string
	}

	var dataSays []Says = []Says{
		{
			Name:   "ucapan",
			Params: "selamat datang",
		},
		{
			Name:   "salam",
			Params: "hallo apa kabar?",
		},
	}

	for _, s := range dataSays {
		b.Run(s.Name, func(b *testing.B) {

			for i := 0; i < b.N; i++ {
				helper.Say(s.Params)
			}
		})
	}
}

/*

jadi kita juga bisa membuat table yang isinya data dan kita dapat gunakan data tersebut
untuk membuat banyak sub benchmark dalam konsep table banchmark

jadi itulah banchmak, pengujian terhadap performa code program kita

jadi selesai sudah materi unit testnya
mudah muahan kalian paham ya

dan selanjutnya kita akan bahasa mengenai materi yang lebih lanjut lagi
yaitu goroutine

*/
